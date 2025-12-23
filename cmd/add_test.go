package cmd

import (
	"fmt"

	"github.com/emmajones/goboot/internal/generator"
	"github.com/spf13/cobra"
)

var addTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Add testing scaffolds and dependencies to the project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := loadProjectContext()
		if err != nil {
			return err
		}

		p, err := generator.Get("test")
		if err != nil {
			return err
		}

		fmt.Printf("Adding testing scaffolds to project %s...\n", ctx.Config.Name)
		return p.Generate(ctx)
	},
}

func init() {
	addCmd.AddCommand(addTestCmd)
}
