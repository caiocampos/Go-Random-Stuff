package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type config struct {
	App app
	DB  database `toml:"database"`
}

type app struct {
	Port int
}

type database struct {
	Server   string
	Database string
}

// TOMLConfig represents the toml configuration file
var TOMLConfig *config

// ReadTOML method reads and parses the TOML configuration file
func ReadTOML() {
	if TOMLConfig == nil {
		TOMLConfig = &config{}
		if _, err := toml.DecodeFile("config.toml", &TOMLConfig); err != nil {
			log.Fatal(err)
		}
	}
}
