package cmd

import (
	"context"
	"fmt"

	corecmd "github.com/wbreza/azd/core/cmd"
)

type listAiModelsCommand struct {
}

func newListAiModelsCommand() corecmd.Command {
	return &listAiModelsCommand{}
}

func (cmd *listAiModelsCommand) Run(ctx context.Context, args []string) (*corecmd.CommandResult, error) {
	fmt.Println("list ai models command")
	return nil, nil
}

type createAiModelCommand struct {
}

func newCreateAiModelCommand() corecmd.Command {
	return &createAiModelCommand{}
}

func (cmd *createAiModelCommand) Run(ctx context.Context, args []string) (*corecmd.CommandResult, error) {
	fmt.Println("create ai model command")
	return nil, nil
}
