package terraform

import (
	"github.com/wbreza/azd/ext"
	"github.com/wbreza/azd/terraform/infra"
	"github.com/wbreza/container/v4"
)

type TerraformExtension struct {
}

func NewTerraformExtension() *TerraformExtension {
	return &TerraformExtension{}
}

func (be *TerraformExtension) Name() string {
	return "terraform"
}

func (be *TerraformExtension) ConfigureContainer(container *container.Container) error {
	return nil
}

func (be *TerraformExtension) Configure(provider *ext.ExtensionProvider) error {
	provider.RegisterInfraProvider("terraform", infra.NewProvider)

	return nil
}

var _ ext.Extension = &TerraformExtension{}
