package main

import (
	"flag"
	"log"
	"time"

	"github.com/meeron/honey-badger/bench"
	"github.com/meeron/honey-badger/config"
	"github.com/meeron/honey-badger/db"
	"github.com/meeron/honey-badger/logger"
	"github.com/meeron/honey-badger/server"
)

var (
	configPath  string
	runBench    bool
	benchTarget string
)

func main() {
	flag.StringVar(&configPath, "config", "", "-config <path_to_config_file>")
	flag.BoolVar(&runBench, "bench", false, "-bench")
	flag.StringVar(&benchTarget, "bench.target", "", "-bench.target 127.0.0.1:18950")
	flag.Parse()

	if benchTarget == "" {
		benchTarget = "127.0.0.1:18950"

		if err := config.Init(configPath); err != nil {
			log.Fatal(err)
		}

		if err := logger.Init(); err != nil {
			log.Fatal(err)
		}

		if err := db.Init(); err != nil {
			log.Fatal(err)
		}

		server.Start()
	}

	if runBench {
		time.Sleep(100 * time.Microsecond)

		bench.Run(benchTarget)

		server.Stop()
	}

	server.Wait()
}
