package cmd

import (
	"github.com/spf13/cobra"
	corecmd "github.com/wbreza/azd/core/cmd"
)

func newAiCommandMetadata() *corecmd.CommandMetadata {
	return &corecmd.CommandMetadata{
		Cobra: &cobra.Command{
			Use:   "ai",
			Short: "Manage AI resources",
		},
		Commands: []corecmd.CommandMetadata{
			{
				Cobra: &cobra.Command{
					Use:   "model",
					Short: "Manage AI models",
				},
				Commands: []corecmd.CommandMetadata{
					{
						Cobra: &cobra.Command{
							Use:   "create",
							Short: "Create new AI model",
						},
						Resolver: newCreateAiModelCommand,
					},
					{
						Cobra: &cobra.Command{
							Use:   "list",
							Short: "List available AI models",
						},
						Resolver: newListAiModelsCommand,
					},
				},
			},
			{
				Cobra: &cobra.Command{
					Use:   "project",
					Short: "Manage AI projects",
				},
				Commands: []corecmd.CommandMetadata{
					{
						Cobra: &cobra.Command{
							Use:   "create",
							Short: "Create new AI project",
						},
						Resolver: newCreateAiProjectCommand,
					},
					{
						Cobra: &cobra.Command{
							Use:   "list",
							Short: "List AI projects",
						},
						Resolver: newListAiProjectsCommand,
					},
				},
			},
		},
	}
}
