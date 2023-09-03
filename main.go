package main

import (
	"log"
	"os"

	"github.com/meeron/honey-badger/config"
	"github.com/meeron/honey-badger/db"
	"github.com/meeron/honey-badger/server"
)

func main() {
	wd, _ := os.Getwd()
	log.Printf("Working dir: %s", wd)

	if err := config.Init("config.json"); err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
