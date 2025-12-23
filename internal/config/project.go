package config

// ProjectType represents the type of project to scaffold.
type ProjectType string

const (
	TypeREST   ProjectType = "rest"
	TypeGRPC   ProjectType = "grpc"
	TypeCLI    ProjectType = "cli"
	TypeWorker ProjectType = "worker"
)

// DBType represents the type of database to include.
type DBType string

const (
	DBPostgres DBType = "postgres"
	DBMySQL    DBType = "mysql"
	DBMongo    DBType = "mongo"
	DBNone     DBType = "none"
)

// AuthType represents the type of authentication to include.
type AuthType string

const (
	AuthJWT    AuthType = "jwt"
	AuthAPIKey AuthType = "apikey"
	AuthNone   AuthType = "none"
)

// ProjectConfig holds the configuration for a new project.
type ProjectConfig struct {
	Name          string      `mapstructure:"name"`
	Module        string      `mapstructure:"module"`
	Type          ProjectType `mapstructure:"type"`
	GoVersion     string      `mapstructure:"go-version"`
	Docker        bool        `mapstructure:"docker"`
	CI            bool        `mapstructure:"ci"`
	DB            DBType      `mapstructure:"db"`
	Auth          AuthType    `mapstructure:"auth"`
	Observability bool        `mapstructure:"observability"`
	DryRun        bool        `mapstructure:"dry-run"`
	Output        string      `mapstructure:"output"`
}
