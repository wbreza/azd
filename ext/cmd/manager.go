package cmd

import (
	corecmd "github.com/wbreza/azd/core/cmd"
)

type Manager interface {
	AddPlugin(name string, resolver any) error
}

type Provider interface {
	RegisterCommand(metadata *corecmd.CommandMetadata) error
}

type Plugin interface {
	Configure(manager Provider) error
}
