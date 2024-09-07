package terraform

import (
	"github.com/wbreza/azd/ext"
	"github.com/wbreza/azd/terraform/infra"
)

type TerraformExtension struct {
}

func NewTerraformExtension() *TerraformExtension {
	return &TerraformExtension{}
}

func (be *TerraformExtension) Name() string {
	return "terraform"
}

func (be *TerraformExtension) Configure(provider *ext.ExtensionProvider) error {
	provider.RegisterInfraProvider("terraform", infra.NewProvider)

	return nil
}

var _ ext.Extension = &TerraformExtension{}
