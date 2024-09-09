package cmd

import (
	extcmd "github.com/wbreza/azd/ext/cmd"
)

type DefaultCommandsPlugin struct {
}

func NewDefaultCommandsPlugin() extcmd.Plugin {
	return &DefaultCommandsPlugin{}
}

func (dp *DefaultCommandsPlugin) Configure(construct extcmd.Provider) error {
	construct.RegisterCommandGroup(newConfigCommandGroupMetadata())
	construct.RegisterCommand(newProvisionCommandMetadata())

	return nil
}

var _ extcmd.Plugin = &DefaultCommandsPlugin{}
