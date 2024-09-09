package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	corecmd "github.com/wbreza/azd/core/cmd"
)

func newConfigCommandMetadata() *corecmd.CommandMetadata {
	return &corecmd.CommandMetadata{
		Cobra: &cobra.Command{
			Use:   "config",
			Short: "Manage azd configuration",
		},
		Commands: []corecmd.CommandMetadata{
			{
				Cobra: &cobra.Command{
					Use:   "show",
					Short: "Show the current configuration",
				},
				Resolver: newShowCommand,
			},
			{
				Cobra: &cobra.Command{
					Use:   "set",
					Short: "Set a configuration value",
				},
				Resolver: newSetCommand,
			},
			{
				Cobra: &cobra.Command{
					Use:   "unset",
					Short: "Unset a configuration value",
				},
				Resolver: newUnsetCommand,
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
