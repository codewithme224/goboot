package context

import (
	"github.com/codewithme224/goboot/internal/config"
)

// ProjectContext holds information about the project being generated or modified.
type ProjectContext struct {
	Config *config.ProjectConfig
	// RootDir is the root directory of the project.
	RootDir string
}
