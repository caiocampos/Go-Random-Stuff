package config

type ConfigLoader interface {
	Load() error
	Get() *Config
}

type Config struct {
	App   app
	DB    database `toml:"database"`
	Cache cache
}

type app struct {
	Port          int
	JsonProcessor string `toml:"json_processor"`
}

type database struct {
	Server   string
	Database string
}

type cache struct {
	Server   string
	Username *string
	Password *string
	Type     string
}
