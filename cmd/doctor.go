package cmd

import (
	"fmt"
	"os"

	"github.com/emmajones/goboot/internal/doctor"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check the health of a goboot project",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		fmt.Printf("🩺 Running diagnostics for project at %s...\n", cwd)

		diag, err := doctor.Check(cwd)
		if err != nil {
			return err
		}

		if len(diag.Issues) == 0 && len(diag.Warnings) == 0 {
			fmt.Println("✅ Everything looks good!")
			return nil
		}

		if len(diag.Issues) > 0 {
			fmt.Println("\n❌ Issues found:")
			for _, issue := range diag.Issues {
				fmt.Printf("  - %s\n", issue)
			}
		}

		if len(diag.Warnings) > 0 {
			fmt.Println("\n⚠️  Warnings:")
			for _, warning := range diag.Warnings {
				fmt.Printf("  - %s\n", warning)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
