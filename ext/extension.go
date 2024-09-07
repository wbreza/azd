package ext

import (
	"github.com/wbreza/azd/ext/cmd"
	"github.com/wbreza/container/v4"
)

type Extension interface {
	Name() string
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

func (ep *ExtensionProvider) RegisterInfraProvider(name string, providerResolver interface{}) error {
	return ep.container.RegisterNamedTransient(name, providerResolver)
}

func (ep *ExtensionProvider) RegisterCommandProvider(name string, providerResolver interface{}) error {
	return ep.commandManager.AddPlugin(name, providerResolver)
}
