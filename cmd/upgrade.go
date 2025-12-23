package cmd

import (
	"fmt"
	"os"

	"github.com/emmajones/goboot/internal/upgrader"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Check for project template upgrades",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		fmt.Printf("🚀 Checking for upgrades at %s...\n", cwd)

		msg, err := upgrader.CheckUpgrade(cwd)
		if err != nil {
			return err
		}

		fmt.Println(msg)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
