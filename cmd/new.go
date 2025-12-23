package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codewithme224/goboot/internal/config"
	"github.com/codewithme224/goboot/internal/generator"
	"github.com/codewithme224/goboot/internal/validator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Scaffold a new Go project",
	Long: `The new command creates a new Go project based on the provided flags.
It validates the input and prepares the project structure.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var cfg config.ProjectConfig

		// Unmarshal flags into config struct using Viper
		if err := viper.Unmarshal(&cfg); err != nil {
			return err
		}

		interactive, _ := cmd.Flags().GetBool("interactive")
		if interactive {
			if err := runInteractive(&cfg); err != nil {
				return err
			}
		}

		// Validate configuration
		if err := validator.ValidateProjectConfig(&cfg); err != nil {
			return err
		}

		// Initialize generator with stdout and dry-run flag
		gen := generator.NewGenerator(os.Stdout, cfg.DryRun)

		// Run generator
		return gen.Generate(&cfg)
	},
}

func runInteractive(cfg *config.ProjectConfig) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🚀 Welcome to goboot Interactive Mode!")
	fmt.Println("--------------------------------------")

	if cfg.Name == "" {
		fmt.Print("Project Name: ")
		name, _ := reader.ReadString('\n')
		cfg.Name = strings.TrimSpace(name)
	}

	if cfg.Module == "" {
		fmt.Print("Go Module Path (e.g. github.com/user/repo): ")
		module, _ := reader.ReadString('\n')
		cfg.Module = strings.TrimSpace(module)
	}

	fmt.Print("Project Type (rest/grpc) [rest]: ")
	pType, _ := reader.ReadString('\n')
	pType = strings.TrimSpace(pType)
	if pType != "" {
		cfg.Type = config.ProjectType(pType)
	}

	fmt.Print("Include Dockerfile? (y/n) [n]: ")
	docker, _ := reader.ReadString('\n')
	cfg.Docker = strings.ToLower(strings.TrimSpace(docker)) == "y"

	fmt.Print("Database Type (postgres/mysql/mongo/none) [none]: ")
	db, _ := reader.ReadString('\n')
	db = strings.TrimSpace(db)
	if db != "" {
		cfg.DB = config.DBType(db)
	}

	fmt.Println("--------------------------------------")
	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Define flags for 'new' command
	newCmd.Flags().StringP("name", "n", "", "Project name")
	newCmd.Flags().StringP("module", "m", "", "Go module path")
	newCmd.Flags().StringP("type", "t", "rest", "Project type (rest | grpc | cli | worker)")
	newCmd.Flags().String("go-version", "1.22", "Go version to use")
	newCmd.Flags().Bool("docker", false, "Include Dockerfile")
	newCmd.Flags().Bool("ci", false, "Include CI/CD workflow")
	newCmd.Flags().String("db", "none", "Database type (postgres | mysql | mongo | none)")
	newCmd.Flags().String("auth", "none", "Authentication type (jwt | apikey | none)")
	newCmd.Flags().Bool("observability", false, "Include observability (metrics, tracing)")
	newCmd.Flags().Bool("dry-run", false, "Run without creating any files")
	newCmd.Flags().StringP("output", "o", ".", "Output directory")
	newCmd.Flags().BoolP("interactive", "i", false, "Run in interactive mode")

	// Bind flags to Viper
	viper.BindPFlags(newCmd.Flags())
}
