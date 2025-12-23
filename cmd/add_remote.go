package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/emmajones/goboot/internal/filesystem"
	"github.com/emmajones/goboot/internal/generator"
	"github.com/emmajones/goboot/internal/template"
	"github.com/spf13/cobra"
)

var remoteUrl string

var addRemoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Add a feature from a remote Git repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := loadProjectContext()
		if err != nil {
			return err
		}

		fmt.Printf("Cloning remote plugin from %s...\n", remoteUrl)

		tempDir, err := os.MkdirTemp("", "goboot-remote-*")
		if err != nil {
			return err
		}
		defer os.RemoveAll(tempDir)

		cloneCmd := exec.Command("git", "clone", "--depth", "1", remoteUrl, tempDir)
		if err := cloneCmd.Run(); err != nil {
			return fmt.Errorf("failed to clone repository: %w", err)
		}

		// Look for templates directory in the cloned repo
		templateDir := filepath.Join(tempDir, "templates")
		if _, err := os.Stat(templateDir); os.IsNotExist(err) {
			// If no templates dir, assume the root is the template dir
			templateDir = tempDir
		}

		fmt.Println("Applying remote templates...")

		// Use a temporary generator to apply the local templates
		writer := filesystem.NewWriter(false)
		renderer := template.NewRenderer()
		base := generator.NewBaseGenerator(writer, renderer)

		// We need to pass the relative path within the DirFS
		// Since we are using the root of the DirFS, templateRoot is "."
		return base.GenerateFromTemplates(ctx.Config, os.DirFS(templateDir), ".", "")
	},
}

func init() {
	addCmd.AddCommand(addRemoteCmd)
	addRemoteCmd.Flags().StringVarP(&remoteUrl, "url", "u", "", "Git repository URL")
	addRemoteCmd.MarkFlagRequired("url")
}
