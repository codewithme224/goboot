package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a feature to an existing project",
	Long:  `The add command allows you to incrementally add features like databases, auth, and gateways to an existing goboot project.`,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
