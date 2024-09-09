package cmd

import (
	corecmd "github.com/wbreza/azd/core/cmd"
	extcmd "github.com/wbreza/azd/ext/cmd"
)

type DefaultCommandsPlugin struct {
}

func NewDefaultCommandsPlugin() extcmd.Plugin {
	return &DefaultCommandsPlugin{}
}

func (dp *DefaultCommandsPlugin) Configure(construct extcmd.Provider) error {
	commands := []*corecmd.CommandMetadata{
		newConfigCommandMetadata(),
		newProvisionCommandMetadata(),
	}

	for _, command := range commands {
		if err := construct.RegisterCommand(command); err != nil {
			return err
		}
	}

	return nil
}

var _ extcmd.Plugin = &DefaultCommandsPlugin{}
