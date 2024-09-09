package ai

import (
	"github.com/wbreza/azd/ai/internal/cmd"
	"github.com/wbreza/azd/ext"
	"github.com/wbreza/container/v4"
)

type AiExtension struct {
}

func NewAiExtension() *AiExtension {
	return &AiExtension{}
}

func (ai *AiExtension) Name() string {
	return "terraform"
}

func (ai *AiExtension) ConfigureContainer(container *container.Container) error {
	return nil
}

// Configure registers the ai commands with the extension provider
func (ai *AiExtension) Configure(provider *ext.ExtensionProvider) error {
	provider.RegisterCommandProvider("ai", cmd.NewAiCommandsPlugin)

	return nil
}

var _ ext.Extension = &AiExtension{}
