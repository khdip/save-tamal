package main

import (
	"log"
	"save-tamal/tamal/storage/postgres"
)

func main() {
	if err := postgres.Migrate(); err != nil {
		log.Fatal(err)
	}
}
