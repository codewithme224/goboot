package validator

import (
	"fmt"
	"regexp"

	"github.com/codewithme224/goboot/internal/config"
)

var (
	// Simple regex for go module path validation
	moduleRegex = regexp.MustCompile(`^[a-z0-9\.\-_/]+$`)
	// Simple regex for project name validation
	nameRegex = regexp.MustCompile(`^[a-z0-9\-_]+$`)
)

// ValidateProjectConfig validates the project configuration.
func ValidateProjectConfig(cfg *config.ProjectConfig) error {
	if cfg.Name == "" {
		return fmt.Errorf("project name is required")
	}
	if !nameRegex.MatchString(cfg.Name) {
		return fmt.Errorf("invalid project name: %s (only lowercase, numbers, hyphens and underscores allowed)", cfg.Name)
	}

	if cfg.Module == "" {
		return fmt.Errorf("go module path is required")
	}
	if !moduleRegex.MatchString(cfg.Module) {
		return fmt.Errorf("invalid module path: %s", cfg.Module)
	}

	if err := validateType(cfg.Type); err != nil {
		return err
	}

	if err := validateDB(cfg.DB); err != nil {
		return err
	}

	if err := validateAuth(cfg.Auth); err != nil {
		return err
	}

	return nil
}

func validateType(t config.ProjectType) error {
	switch t {
	case config.TypeREST, config.TypeGRPC, config.TypeCLI, config.TypeWorker:
		return nil
	default:
		return fmt.Errorf("invalid project type: %s (must be rest, grpc, cli, or worker)", t)
	}
}

func validateDB(db config.DBType) error {
	switch db {
	case config.DBPostgres, config.DBMySQL, config.DBMongo, config.DBNone:
		return nil
	default:
		return fmt.Errorf("invalid database type: %s (must be postgres, mysql, mongo, or none)", db)
	}
}

func validateAuth(auth config.AuthType) error {
	switch auth {
	case config.AuthJWT, config.AuthAPIKey, config.AuthNone:
		return nil
	default:
		return fmt.Errorf("invalid auth type: %s (must be jwt, apikey, or none)", auth)
	}
}
