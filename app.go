package main

import (
	"github.com/daltonclaybrook/swerve/middle"
	"github.com/daltonclaybrook/swerve/server"
)

func main() {
	server := server.NewServer()
	server.AddMiddleware(middle.CORS{Origin: "*", Methods: "POST, GET, OPTIONS", Headers: "*"})
	server.AddControl(&User{})
	server.AddControl(&Asset{})
	server.Start()
}
