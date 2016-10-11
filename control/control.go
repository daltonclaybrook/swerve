package control

import (
	"github.com/daltonclaybrook/swerve/middle"
)

// Route describes an endpoint.
type Route struct {
	Path     string
	Handlers []Handler
}

// Handler describes functions mapped to http methods.
type Handler struct {
	Method  string
	Handler middle.ContextFunc
}

// Control handles routes.
type Control interface {
	Routes() []Route
}
