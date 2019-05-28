package main

import (
	"github.com/go_docker/Go/app"
	"github.com/go_docker/Go/config"
)

func main() {
	config := config.GetConfig()
	port := ":3000"

	app := &app.App{}
	app.Initialize(config)
	app.Run(port)

}
