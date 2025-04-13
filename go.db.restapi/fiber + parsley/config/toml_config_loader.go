package config

import (
	"github.com/BurntSushi/toml"
)

type tomlConfigloader struct {
	config  Config
	started bool
}

func NewTOMLConfigloader() ConfigLoader {
	return &tomlConfigloader{started: false}
}

func (c *tomlConfigloader) Load() error {
	if !c.started {
		TOMLConfig := Config{}
		_, err := toml.DecodeFile("config.toml", &TOMLConfig)
		if err != nil {
			return err
		}
		c.config = TOMLConfig
		c.started = true
	}
	return nil
}

func (c *tomlConfigloader) Get() *Config {
	return &c.config
}
