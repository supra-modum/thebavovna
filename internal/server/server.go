package server

import (
	"bavovna/internal/conf"
	"bavovna/internal/database"
	"bavovna/internal/store"
)

func Start(cfg conf.Config) {
	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()

	// Start listening and serving requests
	router.Run(":5000")
}
