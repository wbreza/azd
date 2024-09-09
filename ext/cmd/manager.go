package cmd

import (
	corecmd "github.com/wbreza/azd/core/cmd"
)

// Manager is an interface that all command managers must implement
type Manager interface {
	AddPlugin(name string, resolver any) error
}

// Provider is an interface that all command providers must implement
type Provider interface {
	RegisterCommand(metadata *corecmd.CommandMetadata) error
}

// Plugin is an interface that all command plugins must implement
type Plugin interface {
	// Configure registers the plugin with the command manager
	Configure(manager Provider) error
}
