package main

import (
	"context"

	"github.com/fatih/color"
	"github.com/wbreza/azd/ai"
	"github.com/wbreza/azd/bicep"
	"github.com/wbreza/azd/cli/internal"
	"github.com/wbreza/azd/cli/internal/cmd"
	"github.com/wbreza/azd/core/ioc"
	"github.com/wbreza/azd/ext"
	extcmd "github.com/wbreza/azd/ext/cmd"
	"github.com/wbreza/azd/terraform"
	"github.com/wbreza/container/v4"
)

func main() {
	ctx := context.Background()
	rootContainer := initContainer()

	if err := initExtensions(ctx, rootContainer); err != nil {
		panic(err)
	}

	if err := runDefaultCommand(ctx, rootContainer); err != nil {
		color.Red("ERROR: %v", err)
	}
}

func initContainer() *container.Container {
	rootContainer := container.New()

	container.MustRegisterInstance(rootContainer, rootContainer)
	container.MustRegisterInstanceAs[ioc.ServiceLocator](rootContainer, rootContainer)
	container.MustRegisterSingleton(rootContainer, ext.NewExtensionProvider)
	container.MustRegisterSingleton(rootContainer, cmd.NewManager)

	container.MustRegisterSingleton(rootContainer, func(commandManager *cmd.Manager) extcmd.Manager {
		return commandManager
	})

	container.MustRegisterSingleton(rootContainer, func(commandManager *cmd.Manager) extcmd.Provider {
		return commandManager
	})

	return rootContainer
}

func initExtensions(ctx context.Context, rootContainer *container.Container) error {
	// Declare the extensions that should be loaded into `azd`
	allExtensions := []ext.Extension{
		internal.NewDefaultExtension(),
		ai.NewAiExtension(),
		bicep.NewBicepExtension(),
		terraform.NewTerraformExtension(),
	}

	return rootContainer.Call(ctx, func(extensionProvider *ext.ExtensionProvider) error {
		for _, extension := range allExtensions {
			// Each extension could register its own components into the container
			if err := extension.ConfigureContainer(rootContainer); err != nil {
				return err
			}

			// Extension has a configuration entrypoint that allows it to extensible components
			if err := extension.Configure(extensionProvider); err != nil {
				return err
			}
		}

		return nil
	})
}

// Run the default azd command
func runDefaultCommand(ctx context.Context, rootContainer *container.Container) error {
	return rootContainer.Call(ctx, func(commandManager *cmd.Manager) error {
		if err := commandManager.Initialize(ctx); err != nil {
			return err
		}

		if err := commandManager.Run(ctx); err != nil {
			return err
		}

		return nil
	})
}
