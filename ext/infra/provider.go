package infra

import "context"

// Provider is the interface that all infrastructure providers must implement
type Provider interface {
	// Name returns the name of the provider
	Name() string
	// Initialize initializes the provider with the project path
	Initialize(ctx context.Context, projectPath string) error
	// Plan plans the infrastructure changes
	State(ctx context.Context) error
	// Apply applies the infrastructure changes
	Deploy(ctx context.Context) error
	// Preview previews the infrastructure changes
	Preview(ctx context.Context) error
	// Destroy destroys the infrastructure
	Destroy(ctx context.Context) error
}
