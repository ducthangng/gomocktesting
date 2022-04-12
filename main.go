package main

import (
	"gomockt/config"
	"gomockt/model/db/sqlcon"
	"log"
)

func main() {
	config, err := config.ReadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	sqlcon.OpenConnection(config)
}
