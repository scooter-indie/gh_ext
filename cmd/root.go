package cmd

import (
	"github.com/spf13/cobra"
)

var version = "dev"

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gh-pm",
		Short:   "GitHub Project Management CLI",
		Long:    `gh-pm is a unified GitHub CLI extension for project management, sub-issue hierarchy, and project templating.`,
		Version: version,
	}

	return cmd
}

func Execute() error {
	return NewRootCommand().Execute()
}
