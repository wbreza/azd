package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

type CommandMetadata struct {
	Cobra    *cobra.Command
	Resolver any
}

type CommandGroupMetadata struct {
	CommandMetadata

	Commands []CommandMetadata
}

type Command interface {
	Run(ctx context.Context, args []string) (*CommandResult, error)
}

type CommandResult struct {
}
