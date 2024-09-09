package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	corecmd "github.com/wbreza/azd/core/cmd"
	extcmd "github.com/wbreza/azd/ext/cmd"
	"github.com/wbreza/container/v4"
)

type Manager struct {
	container        *container.Container
	rootCobraCommand *cobra.Command
	plugins          []string
}

func NewManager(rootContainer *container.Container) *Manager {
	return &Manager{
		rootCobraCommand: &cobra.Command{
			Use:   "azd",
			Short: "Azure Developer CLI",
		},
		container: rootContainer,
		plugins:   []string{},
	}
}

func (m *Manager) Initialize(ctx context.Context) error {
	for _, pluginName := range m.plugins {
		var plugin extcmd.Plugin
		if err := m.container.ResolveNamed(ctx, pluginName, &plugin); err != nil {
			return err
		}

		if err := plugin.Configure(m); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) AddPlugin(name string, resolver any) error {
	pluginName := fmt.Sprintf("%s-commands-plugin", name)
	if err := m.container.RegisterNamedSingleton(pluginName, resolver); err != nil {
		return err
	}

	m.plugins = append(m.plugins, pluginName)

	return nil
}

func (m *Manager) RegisterCommand(metadata *corecmd.CommandMetadata) error {
	_, err := m.registerCommand(m.rootCobraCommand, metadata)
	return err
}

func (m *Manager) RegisterCommandGroup(metadata *corecmd.CommandGroupMetadata) error {
	commandGroup, err := m.registerCommand(m.rootCobraCommand, &metadata.CommandMetadata)
	if err != nil {
		return err
	}

	for _, commandMetadata := range metadata.Commands {
		m.registerCommand(commandGroup, &commandMetadata)
	}

	return nil
}

func (m *Manager) Run(ctx context.Context) error {
	return m.rootCobraCommand.ExecuteContext(ctx)
}

func (m *Manager) registerCommand(parent *cobra.Command, commandMetadata *corecmd.CommandMetadata) (*cobra.Command, error) {
	command := commandMetadata.Cobra

	if commandMetadata.Resolver != nil {
		commandName := fmt.Sprintf("%s-%s", parent.CommandPath(), command.Name())
		if err := m.container.RegisterNamedTransient(commandName, commandMetadata.Resolver); err != nil {
			return nil, err
		}

		command.RunE = func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			var command corecmd.Command
			if err := m.container.ResolveNamed(ctx, commandName, &command); err != nil {
				return err
			}

			_, err := command.Run(ctx, args)
			if err != nil {
				return err
			}

			return nil
		}
	}

	parent.AddCommand(command)

	return command, nil
}

var _ extcmd.Manager = &Manager{}
