package cmd

import (
	"context"
	"fmt"

	corecmd "github.com/wbreza/azd/core/cmd"
)

func newConfigCommandGroup() *corecmd.CommandGroupMetadata {
	return &corecmd.CommandGroupMetadata{
		CommandMetadata: corecmd.CommandMetadata{
			Name:        "config",
			Short:       "config",
			Description: "Manage azd configuration",
		},
		Commands: []corecmd.CommandMetadata{
			{
				Name:        "show",
				Short:       "show",
				Description: "Show the current configuration",
				Resolver:    newShowCommand,
			},
			{
				Name:        "set",
				Short:       "set",
				Description: "Set a configuration value",
				Resolver:    newSetCommand,
			},
			{
				Name:        "unset",
				Short:       "unset",
				Description: "Unset a configuration value",
				Resolver:    newUnsetCommand,
			},
		},
	}
}

type showCommand struct {
}

func newShowCommand() corecmd.Command {
	return &showCommand{}
}

func (sc *showCommand) Run(ctx context.Context, args []string) (*corecmd.CommandResult, error) {
	fmt.Println("show command")
	return nil, nil
}

type setCommand struct {
}

func newSetCommand() corecmd.Command {
	return &setCommand{}
}

func (sc *setCommand) Run(ctx context.Context, args []string) (*corecmd.CommandResult, error) {
	fmt.Println("set command")
	return nil, nil
}

type unsetCommand struct {
}

func newUnsetCommand() corecmd.Command {
	return &unsetCommand{}
}

func (sc *unsetCommand) Run(ctx context.Context, args []string) (*corecmd.CommandResult, error) {
	fmt.Println("unset command")
	return nil, nil
}
