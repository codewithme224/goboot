package cmd

import (
	"fmt"

	"github.com/codewithme224/goboot/internal/generator"
	"github.com/spf13/cobra"
)

var addK8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Add Kubernetes manifests to the project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := loadProjectContext()
		if err != nil {
			return err
		}

		p, err := generator.Get("k8s")
		if err != nil {
			return err
		}

		fmt.Printf("Adding Kubernetes manifests to project %s...\n", ctx.Config.Name)
		return p.Generate(ctx)
	},
}

func init() {
	addCmd.AddCommand(addK8sCmd)
}
