package cmd

import "context"

type CommandMetadata struct {
	Name        string
	Short       string
	Description string
	Aliases     []string
	Resolver    any
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
