package bicep

import (
	"github.com/wbreza/azd/bicep/infra"
	"github.com/wbreza/azd/ext"
)

var Extensions = map[string]ext.Extension{}

func init() {
	Extensions["bicep"] = NewBicepExtension()
}

type BicepExtension struct {
}

func NewBicepExtension() ext.Extension {
	return &BicepExtension{}
}

func (be *BicepExtension) Name() string {
	return "bicep"
}

func (be *BicepExtension) Configure(provider *ext.ExtensionProvider) error {
	provider.RegisterInfraProvider("bicep", infra.NewProvider)

	return nil
}
