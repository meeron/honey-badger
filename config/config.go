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
	Port uint16
}

type BadgerConfig struct {
	DataDirPath string
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
		return nil
	}
	defer f.Close()

	fileConfing := Config{}
	decoder := json.NewDecoder(f)

	if err := decoder.Decode(&fileConfing); err != nil {
		return err
	}

	if fileConfing.Server.Port <= 1023 {
		log.Printf("Invalid config value: 'Server.Port'. Using default value '%d'", defaultConfig.Server.Port)
		fileConfing.Server.Port = defaultConfig.Server.Port
	}

	if fileConfing.Badger.DataDirPath == "" {
		log.Printf("Invalid config value: 'Badger.DataDirPath'. Using default value '%s'", defaultConfig.Badger.DataDirPath)
		fileConfing.Badger.DataDirPath = current.Badger.DataDirPath
	}

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
		},
		Server: ServerConfig{
			Port: 18950,
		},
	}
}
