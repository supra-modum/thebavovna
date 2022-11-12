package config

import (
	"github.com/go-pg/pg/v10"
	"log"
	"os"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("BAV_DB_USER"),
		Password: os.Getenv("BAV_DB_PASSWORD"),
		Addr:     os.Getenv("BAV_DB_HOST") + ":" + os.Getenv("BAV_DB_PORT"),
		Database: os.Getenv("BAV_DB_NAME"),
	}
	
	var db = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to db")

	return db
}
