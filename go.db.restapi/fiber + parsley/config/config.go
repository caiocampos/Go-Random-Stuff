package config

type ConfigLoader interface {
	Load() error
	Get() *Config
}

type Config struct {
	App app
	DB  database `toml:"database"`
}

type app struct {
	Port          int
	JsonProcessor string `toml:"json_processor"`
}

type database struct {
	Server   string
	Database string
}
