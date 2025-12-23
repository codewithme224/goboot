package cmd

import (
	"fmt"

	"github.com/emmajones/goboot/internal/generator"
	"github.com/spf13/cobra"
)

// addGatewayCmd represents the add gateway command
var addGatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Add gRPC-Gateway to the project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := loadProjectContext()
		if err != nil {
			return err
		}

		p, err := generator.Get("gateway")
		if err != nil {
			return err
		}

		fmt.Printf("Adding gRPC-Gateway to project %s...\n", ctx.Config.Name)
		return p.Generate(ctx)
	},
}

func init() {
	addCmd.AddCommand(addGatewayCmd)
}
