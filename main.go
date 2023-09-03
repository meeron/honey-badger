package main

import (
	"flag"
	"log"

	"github.com/meeron/honey-badger/config"
	"github.com/meeron/honey-badger/db"
	"github.com/meeron/honey-badger/logger"
	"github.com/meeron/honey-badger/server"
)

var (
	configPath string
)

func main() {
	flag.StringVar(&configPath, "config", "", "-config <path_to_config_file>")
	flag.Parse()

	if err := config.Init(configPath); err != nil {
		log.Fatal(err)
	}

	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
