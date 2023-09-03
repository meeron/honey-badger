package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Badger BadgerConfig
	Server ServerConfig
	Logger LoggerConfig
}

type ServerConfig struct {
	Port             uint16
	MaxRecvMsgSizeMb int
}

type BadgerConfig struct {
	DataDirPath string
	GCPeriodMin int
}

type LoggerConfig struct {
	Dir string
}

var current Config

func Init(configFilePath string) error {
	defaultConfig := createDefaultConfig()

	if configFilePath == "" {
		current = defaultConfig
		return nil
	}

	f, err := os.OpenFile(configFilePath, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	fileConfing := Config{}
	decoder := json.NewDecoder(f)

	if err := decoder.Decode(&fileConfing); err != nil {
		return err
	}

	validateConfig(&fileConfing, &defaultConfig)

	current = fileConfing

	return nil
}

func Get() Config {
	return current
}

func createDefaultConfig() Config {
	return Config{
		Badger: BadgerConfig{
			DataDirPath: "data",
			GCPeriodMin: 60,
		},
		Server: ServerConfig{
			Port:             18950,
			MaxRecvMsgSizeMb: 100,
		},
		Logger: LoggerConfig{
			Dir: "logs",
		},
	}
}

func validateConfig(config *Config, defaultConfig *Config) {
	if config.Server.Port <= 1023 {
		config.Server.Port = defaultConfig.Server.Port
	}

	if config.Server.MaxRecvMsgSizeMb < 4 {
		config.Server.MaxRecvMsgSizeMb = defaultConfig.Server.MaxRecvMsgSizeMb
	}

	if config.Badger.DataDirPath == "" {
		config.Badger.DataDirPath = defaultConfig.Badger.DataDirPath
	}

	if config.Badger.GCPeriodMin <= 0 {
		config.Badger.GCPeriodMin = defaultConfig.Badger.GCPeriodMin
	}

	if config.Logger.Dir == "" {
		config.Logger.Dir = defaultConfig.Logger.Dir
	}
}
