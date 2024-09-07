package internal

import (
	"github.com/wbreza/azd/cli/internal/cmd"
	"github.com/wbreza/azd/ext"
)

type DefaultExtension struct {
}

func NewDefaultExtension() *DefaultExtension {
	return &DefaultExtension{}
}

func (de *DefaultExtension) Name() string {
	return "default"
}

func (de *DefaultExtension) Configure(provider *ext.ExtensionProvider) error {
	provider.RegisterCommandProvider("default", cmd.NewDefaultCommandsPlugin)

	return nil
}

var _ ext.Extension = &DefaultExtension{}
