package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/wbreza/azd/bicep"
	"github.com/wbreza/azd/core/infra"
	"github.com/wbreza/azd/core/ioc"
	"github.com/wbreza/azd/ext"
	"github.com/wbreza/azd/terraform"
	"github.com/wbreza/container/v4"
)

func main() {
	ctx := context.Background()
	rootContainer := initContainer()
	if err := initExtensions(rootContainer); err != nil {
		panic(err)
	}

	var providerFactory *infra.ProviderFactory
	if err := rootContainer.Resolve(ctx, &providerFactory); err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: azd <provider>")
		os.Exit(1)
	}

	providerName := os.Args[1]
	if err := exec(ctx, providerFactory, providerName); err != nil {
		color.Red("ERROR: %v", err)
	}
}

func exec(ctx context.Context, providerFactory *infra.ProviderFactory, providerName string) error {
	infraProvider, err := providerFactory.Create(ctx, providerName)
	if err != nil {
		return fmt.Errorf("No infra provider found with name '%s'. Details: %w", providerName, err)
	}

	color.Cyan("Created %s provider\n", infraProvider.Name())

	return nil
}

func initContainer() *container.Container {
	rootContainer := container.New()

	container.MustRegisterInstanceAs[ioc.ServiceLocator](rootContainer, rootContainer)
	container.MustRegisterSingleton(rootContainer, infra.NewProviderFactory)

	return rootContainer
}

func initExtensions(rootContainer *container.Container) error {
	extensionProvider := ext.NewExtensionProvider(rootContainer)
	extensions := []map[string]ext.Extension{bicep.Extensions, terraform.Extensions}
	for _, group := range extensions {
		for _, extension := range group {
			if err := extension.Configure(extensionProvider); err != nil {
				return err
			}
		}
	}

	return nil
}
