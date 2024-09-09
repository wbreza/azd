package cmd

import (
	"context"
	"fmt"

	corecmd "github.com/wbreza/azd/core/cmd"
)

type createAiProjectCommand struct {
}

func newCreateAiProjectCommand() corecmd.Command {
	return &createAiProjectCommand{}
}

func (c *createAiProjectCommand) Run(ctx context.Context, args []string) (*corecmd.CommandResult, error) {
	fmt.Println("create ai project command")
	return nil, nil
}

type listAiProjectsCommand struct {
}

func newListAiProjectsCommand() corecmd.Command {
	return &listAiProjectsCommand{}
}

func (c *listAiProjectsCommand) Run(ctx context.Context, args []string) (*corecmd.CommandResult, error) {
	fmt.Println("list ai projects command")
	return nil, nil
}
