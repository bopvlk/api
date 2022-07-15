package api

import "github.com/BO/6.ServerAndDB/storage"

//general instance for API server og REST application
type Config struct {
	//port
	BindAddr string `toml:"bind_addr"`
	// logger level
	LoggerLevel string `toml:"logger_level"`
	// Store config
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":5050",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
