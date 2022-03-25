package config

// Version defines the pbf_demo version.
var Version string

// Config defines the configuration structure.
type Config struct {
	General struct {
		Name        string `mapstructure:"name"`
		LogLevel    int    `mapstructure:"log_level"`
		ReleaseMode bool   `mapstructure:"releaseMode"`
	} `mapstructure:"general"`

	Service struct {
		Grpc int    `mapstructure:"grpc"`
		Http string `mapstructure:"http"`
	} `mapstructure:"serviceport"`

	PostgreSQL struct {
		DriverName  string `mapstructure:"driverName"`
		DSN         string `mapstructure:"dsn"`
		Automigrate bool
	} `mapstructure:"postgresql"`

	Redis struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"redis"`
}

// C holds the global configuration.
var C Config
