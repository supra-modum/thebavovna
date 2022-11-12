package main

import (
	"bavovna/internal/conf"
	"bavovna/internal/server"
)

func main() {
	server.Start(conf.NewConfig())
}
