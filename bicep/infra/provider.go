package infra

import (
	"context"

	"github.com/wbreza/azd/ext/infra"
)

type Provider struct {
}

func NewProvider() infra.Provider {
	return &Provider{}
}

func (p *Provider) Name() string {
	return "bicep"
}

func (p *Provider) Initialize(ctx context.Context, projectPath string) error {
	return nil
}

func (p *Provider) State(ctx context.Context) error {
	return nil
}

func (p *Provider) Deploy(ctx context.Context) error {
	return nil
}

func (p *Provider) Preview(ctx context.Context) error {
	return nil
}

func (p *Provider) Destroy(ctx context.Context) error {
	return nil
}
