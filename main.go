package main

import (
	"flag"
	"log"

	"github.com/meeron/honey-badger/bench"
	"github.com/meeron/honey-badger/config"
	"github.com/meeron/honey-badger/db"
	"github.com/meeron/honey-badger/logger"
	"github.com/meeron/honey-badger/server"
)

var (
	configPath  string
	benchTarget string
)

func main() {
	flag.StringVar(&configPath, "config", "", "-config <path_to_config_file>")
	flag.StringVar(&benchTarget, "bench", "", "-bench 127.0.0.1:18950")
	flag.Parse()

	if benchTarget != "" {
		bench.Run(benchTarget)
		return
	}

	if err := config.Init(configPath); err != nil {
		log.Fatal(err)
	}

	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	server := server.New(config.Get().Server)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
