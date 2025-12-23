package cmd

import (
	"fmt"

	"github.com/emmajones/goboot/internal/generator"
	"github.com/spf13/cobra"
)

var addObservabilityCmd = &cobra.Command{
	Use:   "observability",
	Short: "Add observability scaffolding to the project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := loadProjectContext()
		if err != nil {
			return err
		}

		p, err := generator.Get("observability")
		if err != nil {
			return err
		}

		fmt.Printf("Adding observability to project %s...\n", ctx.Config.Name)
		return p.Generate(ctx)
	},
}

func init() {
	addCmd.AddCommand(addObservabilityCmd)
}
