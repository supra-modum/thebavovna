package server

import "bavovna/internal/config"

func Start() {
	config.Connect()

	router := setRouter()

	// Start listening and serving requests
	router.Run(":5000")
}
