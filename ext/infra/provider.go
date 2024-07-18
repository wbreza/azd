package infra

import "context"

type Provider interface {
	Name() string
	Initialize(ctx context.Context, projectPath string) error
	State(ctx context.Context) error
	Deploy(ctx context.Context) error
	Preview(ctx context.Context) error
	Destroy(ctx context.Context) error
}
