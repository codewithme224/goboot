package cmd

import (
	"fmt"

	"github.com/emmajones/goboot/internal/generator"
	"github.com/spf13/cobra"
)

var addCiCmd = &cobra.Command{
	Use:   "ci",
	Short: "Add CI/CD workflows to the project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := loadProjectContext()
		if err != nil {
			return err
		}

		p, err := generator.Get("ci")
		if err != nil {
			return err
		}

		fmt.Printf("Adding CI/CD workflows to project %s...\n", ctx.Config.Name)
		return p.Generate(ctx)
	},
}

func init() {
	addCmd.AddCommand(addCiCmd)
}
