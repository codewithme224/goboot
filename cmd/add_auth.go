package cmd

import (
	"fmt"

	"github.com/codewithme224/goboot/internal/config"
	"github.com/codewithme224/goboot/internal/generator"
	"github.com/spf13/cobra"
)

var authType string

// addAuthCmd represents the add auth command
var addAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Add authentication to the project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := loadProjectContext()
		if err != nil {
			return err
		}

		ctx.Config.Auth = config.AuthType(authType)

		p, err := generator.Get("auth")
		if err != nil {
			return err
		}

		fmt.Printf("Adding %s auth to project %s...\n", authType, ctx.Config.Name)
		return p.Generate(ctx)
	},
}

func init() {
	addCmd.AddCommand(addAuthCmd)
	addAuthCmd.Flags().StringVarP(&authType, "type", "t", "jwt", "Auth type (jwt)")
	addAuthCmd.MarkFlagRequired("type")
}
