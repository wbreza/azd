package internal

import (
	"github.com/wbreza/azd/cli/internal/cmd"
	coreinfra "github.com/wbreza/azd/core/infra"
	"github.com/wbreza/azd/ext"
	"github.com/wbreza/container/v4"
)

// DefaultExtension is the default extension for azd.
// It provides required registrations for advanced us cases and also registers the default commands.
type DefaultExtension struct {
}

func NewDefaultExtension() *DefaultExtension {
	return &DefaultExtension{}
}

func (de *DefaultExtension) Name() string {
	return "default"
}

// ConfigureContainer registers the required components with the container
func (de *DefaultExtension) ConfigureContainer(c *container.Container) error {
	container.MustRegisterSingleton(c, coreinfra.NewProviderFactory)

	return nil
}

// Configure registers the default commands with the extension provider
func (de *DefaultExtension) Configure(provider *ext.ExtensionProvider) error {
	return provider.RegisterCommandProvider("default", cmd.NewDefaultCommandsPlugin)
}

var _ ext.Extension = &DefaultExtension{}
