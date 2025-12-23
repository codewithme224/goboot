package cmd

import (
	"fmt"

	"github.com/codewithme224/goboot/internal/config"
	"github.com/codewithme224/goboot/internal/generator"
	"github.com/spf13/cobra"
)

var dbType string

// addDbCmd represents the add db command
var addDbCmd = &cobra.Command{
	Use:   "db",
	Short: "Add a database to the project",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := loadProjectContext()
		if err != nil {
			return err
		}

		ctx.Config.DB = config.DBType(dbType)

		p, err := generator.Get("db")
		if err != nil {
			return err
		}

		fmt.Printf("Adding %s database to project %s...\n", dbType, ctx.Config.Name)
		return p.Generate(ctx)
	},
}

func init() {
	addCmd.AddCommand(addDbCmd)
	addDbCmd.Flags().StringVarP(&dbType, "type", "t", "postgres", "Database type (postgres | mysql | mongo)")
	addDbCmd.MarkFlagRequired("type")
}
