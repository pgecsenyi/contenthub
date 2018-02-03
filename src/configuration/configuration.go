package configuration

// Configuration Stores the server configuration.
type Configuration struct {
	ContentDirectoryPath string
	LogFilePath          string
	Port                 int
	UseBuiltInServer     bool
}
