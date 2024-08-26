package constants

// config file slice
const ConfigType = "yaml"
const ConfigDir = "resources/configs"
const ApplicationConfig = "application"
const BusinessConfig = "business"

var ConfigFiles = []string{
	ApplicationConfig,
	BusinessConfig,
}

// business logic constants
const (
	ServerInstanceResourceType = "SERVER_INSTANCE"
)
