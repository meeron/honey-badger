package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Badger BadgerConfig
	Server ServerConfig
}

type ServerConfig struct {
	Port             uint16
	MaxRecvMsgSizeMb int
}

type BadgerConfig struct {
	DataDirPath string
	GCPeriodMin int
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
	}
}

func validateConfig(config *Config, defaultConfig *Config) {
	if config.Server.Port <= 1023 {
		log.Printf("'Server.Port': using default value '%d'", defaultConfig.Server.Port)
		config.Server.Port = defaultConfig.Server.Port
	}

	if config.Server.MaxRecvMsgSizeMb < 4 {
		log.Printf("'Server.MaxRecvMsgSizeMb': using default value '%d'", defaultConfig.Server.MaxRecvMsgSizeMb)
		config.Server.MaxRecvMsgSizeMb = defaultConfig.Server.MaxRecvMsgSizeMb
	}

	if config.Badger.DataDirPath == "" {
		log.Printf("'Badger.DataDirPath': using default value '%s'", defaultConfig.Badger.DataDirPath)
		config.Badger.DataDirPath = defaultConfig.Badger.DataDirPath
	}

	if config.Badger.GCPeriodMin <= 0 {
		log.Printf("'Badger.GCPeriodMin': using default value '%d'", defaultConfig.Badger.GCPeriodMin)
		config.Badger.GCPeriodMin = defaultConfig.Badger.GCPeriodMin
	}
}
