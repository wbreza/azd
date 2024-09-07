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
	if err := container.RegisterNamedSingleton(pluginName, resolver); err != nil {
		return err
	}

	m.plugins = append(m.plugins, pluginName)

	return nil
}

func (m *Manager) RegisterCommand(metadata *corecmd.CommandMetadata) error {
	return nil
}

func (m *Manager) RegisterCommandGroup(metadata *corecmd.CommandGroupMetadata) error {
	commandGroup := &cobra.Command{
		Use:   metadata.Name,
		Short: metadata.Short,
		Long:  metadata.Description,
	}

	m.rootCobraCommand.AddCommand(commandGroup)

	for _, commandMetadata := range metadata.Commands {
		commandName := fmt.Sprintf("%s-%s", commandGroup.CommandPath(), commandMetadata.Name)
		m.container.RegisterNamedTransient(commandName, commandMetadata.Resolver)

		commandGroup.AddCommand(&cobra.Command{
			Use:   commandMetadata.Name,
			Short: commandMetadata.Short,
			Long:  commandMetadata.Description,
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()

				var command corecmd.Command
				if err := m.container.ResolveNamed(cmd.Context(), commandName, &command); err != nil {
					return err
				}

				_, err := command.Run(ctx, args)
				if err != nil {
					return err
				}

				return nil
			},
		})
	}

	return nil
}

func (m *Manager) Run(ctx context.Context) error {
	return m.rootCobraCommand.ExecuteContext(ctx)
}

var _ extcmd.Manager = &Manager{}
