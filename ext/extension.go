package ext

import "github.com/wbreza/container/v4"

type Extension interface {
	Name() string
	Configure(provider *ExtensionProvider) error
}

type ExtensionProvider struct {
	container *container.Container
}

func NewExtensionProvider(rootContainer *container.Container) *ExtensionProvider {
	return &ExtensionProvider{
		container: rootContainer,
	}
}

func (ep *ExtensionProvider) RegisterInfraProvider(name string, provider interface{}) error {
	return ep.container.RegisterNamedTransient(name, provider)
}
