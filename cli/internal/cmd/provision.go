package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	corecmd "github.com/wbreza/azd/core/cmd"
	"github.com/wbreza/azd/core/infra"
)

func newProvisionCommandMetadata() *corecmd.CommandMetadata {
	return &corecmd.CommandMetadata{
		Cobra: &cobra.Command{
			Use:   "provision <name>",
			Short: "Provisions Azure resources",
			Args:  cobra.ExactArgs(1),
		},
		Resolver: newProvisionCommand,
	}
}

type provisionCommand struct {
	providerFactory *infra.ProviderFactory
}

func newProvisionCommand(providerFactory *infra.ProviderFactory) corecmd.Command {
	return &provisionCommand{
		providerFactory: providerFactory,
	}
}

func (pc *provisionCommand) Run(ctx context.Context, args []string) (*corecmd.CommandResult, error) {
	provisionProvider, err := pc.providerFactory.Create(ctx, args[0])
	if err != nil {
		return nil, err
	}

	fmt.Printf("provision command with '%s' provider\n", provisionProvider.Name())
	return nil, nil
}
