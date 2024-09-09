package ext

import (
	"github.com/wbreza/azd/ext/cmd"
	"github.com/wbreza/container/v4"
)

// Extension is an interface that all extensions must implement
type Extension interface {
	// Name returns the name of the extension
	Name() string
	// ConfigureContainer registers the required components with the container
	ConfigureContainer(container *container.Container) error
	// Configure registers the extension with the extension provider
	Configure(provider *ExtensionProvider) error
}

type ExtensionProvider struct {
	container      *container.Container
	commandManager cmd.Manager
}

func NewExtensionProvider(rootContainer *container.Container, commandManager cmd.Manager) *ExtensionProvider {
	return &ExtensionProvider{
		container:      rootContainer,
		commandManager: commandManager,
	}
}

// RegisterInfraProvider registers a new infrastructure provider with the extension provider
func (ep *ExtensionProvider) RegisterInfraProvider(name string, providerResolver interface{}) error {
	return ep.container.RegisterNamedTransient(name, providerResolver)
}

// RegisterCommandProvider registers a new command provider with the extension provider
func (ep *ExtensionProvider) RegisterCommandProvider(name string, providerResolver interface{}) error {
	return ep.commandManager.AddPlugin(name, providerResolver)
}
