package main

import (
	"log"

	"github.com/meeron/honey-badger/api"
	"github.com/meeron/honey-badger/db"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal(err)
	}

	api.Run(":8080")
}
