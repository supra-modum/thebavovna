package main

import (
	"bavovna/internal/cli"
	"bavovna/internal/conf"
	"bavovna/internal/server"
)

func main() {
	cli.Parse()
	server.Start(conf.NewConfig())
}
