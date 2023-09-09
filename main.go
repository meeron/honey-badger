package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/meeron/honey-badger/bench"
	"github.com/meeron/honey-badger/config"
	"github.com/meeron/honey-badger/db"
	"github.com/meeron/honey-badger/logger"
	"github.com/meeron/honey-badger/server"
)

var (
	configPath   string
	benchTarget  string
	printVersion bool
	version      string
	build        string
)

func main() {
	flag.StringVar(&configPath, "config", "", "-config <path_to_config_file>")
	flag.StringVar(&benchTarget, "bench", "", "-bench 127.0.0.1:18950")
	flag.BoolVar(&printVersion, "version", false, "-version")
	flag.Parse()

	if printVersion {
		fmt.Printf("version: %s build: %s\n", version, build)
		return
	}

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

	dbCtx := db.CreateCtx(config.Get().Badger)
	defer dbCtx.Close()

	server := server.New(config.Get().Server, dbCtx)

	if err := dbCtx.LoadDbs(); err != nil {
		log.Fatal(err)
	}

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
