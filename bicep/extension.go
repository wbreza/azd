package bicep

import (
	"github.com/wbreza/azd/bicep/infra"
	"github.com/wbreza/azd/ext"
	"github.com/wbreza/container/v4"
)

type BicepExtension struct {
}

func NewBicepExtension() *BicepExtension {
	return &BicepExtension{}
}

func (be *BicepExtension) Name() string {
	return "bicep"
}

func (be *BicepExtension) ConfigureContainer(container *container.Container) error {
	return nil
}

// Configure registers the bicep provider with the extension provider
func (be *BicepExtension) Configure(provider *ext.ExtensionProvider) error {
	provider.RegisterInfraProvider("bicep", infra.NewProvider)

	return nil
}

var _ ext.Extension = &BicepExtension{}
