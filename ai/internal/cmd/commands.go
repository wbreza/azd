package cmd

import (
	corecmd "github.com/wbreza/azd/core/cmd"
	extcmd "github.com/wbreza/azd/ext/cmd"
)

type AiCommandsPlugin struct {
}

func NewAiCommandsPlugin() extcmd.Plugin {
	return &AiCommandsPlugin{}
}

func (ai *AiCommandsPlugin) Configure(construct extcmd.Provider) error {
	commands := []*corecmd.CommandMetadata{
		newAiCommandMetadata(),
	}

	for _, command := range commands {
		if err := construct.RegisterCommand(command); err != nil {
			return err
		}
	}

	return nil
}

var _ extcmd.Plugin = &AiCommandsPlugin{}
