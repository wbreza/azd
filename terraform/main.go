package terraform

import (
	"github.com/wbreza/azd/ext"
	"github.com/wbreza/azd/terraform/infra"
)

var Extensions = map[string]ext.Extension{}

func init() {
	Extensions["bicep"] = NewTerraformExtension()
}

type TerraformExtension struct {
}

func NewTerraformExtension() ext.Extension {
	return &TerraformExtension{}
}

func (be *TerraformExtension) Name() string {
	return "terraform"
}

func (be *TerraformExtension) Configure(provider *ext.ExtensionProvider) error {
	provider.RegisterInfraProvider("terraform", infra.NewProvider)

	return nil
}
