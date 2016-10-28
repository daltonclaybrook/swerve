package control

import (
	"github.com/daltonclaybrook/swerve/middle"
)

// Route describes an endpoint.
type Route struct {
	// Specifies the url path of the route, such as "/user"
	// The server uses github.com/gorilla/mux for routing,
	// so you can use patterns and variables like "/{id:[0-9]+}"
	Path string

	// The handlers which will be used for this path.
	Handlers []Handler
}

// Handler describes functions mapped to http methods.
type Handler struct {
	// The http method (GET, POST, PUT...) which calls HandlerFunc
	Method string

	// The function which is called when a route is handled.
	HandlerFunc middle.ContextFunc

	// Ordered list of middleware handlers to execute before the route is handled
	Middleware []middle.Handler
}

// Control handles routes.
type Control interface {

	// Returns all routes containing all handlers which
	// should be used for this Control
	Routes() []Route
}
