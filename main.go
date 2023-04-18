package main

import (
	"main/config"
	m "main/middlewares"
	"main/routes"
)

func main() {
	// init database
	config.Init()

	// create a new echo instance
	e := routes.New()

	m.LogMiddleWares(e)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
