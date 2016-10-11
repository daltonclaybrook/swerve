package main

import (
	"github.com/daltonclaybrook/swerve/control"
	"github.com/daltonclaybrook/swerve/middle"
	"github.com/daltonclaybrook/swerve/server"
)

func main() {
	server := server.NewWebServer()
	server.RegisterMiddleware(middle.CORS{Origin: "*", Methods: "POST, GET, OPTIONS", Headers: "*"})
	server.RegisterControl(control.NewTransfer())
	server.Start()
}
