package crud

import (
	"fmt"
	"github.com/daltonclaybrook/swerve/control"
	"github.com/daltonclaybrook/swerve/middle"
	"net/http"
)

const (
	// Create a model.
	Create = iota
	// Find models.
	Find = iota
	// FindOne model.
	FindOne = iota
	// Update a model.
	Update = iota
	// Delete a model.
	Delete = iota
)

// Route is used to express a typical CRUD operation.
type Route struct {

	// One of the operations defined above, e.g.
	// Create, Find, FindOne, Update, and Delete
	Op int

	// The function which is called when the route is handled.
	HandlerFunc middle.ContextFunc

	// An ordered list of middleware handlers
	Middleware []middle.Handler
}

// Handler defines a type which handles all CRUD methods.
type Handler interface {
	Create(w http.ResponseWriter, r *http.Request, c middle.Context)
	Find(w http.ResponseWriter, r *http.Request, c middle.Context)
	FindOne(w http.ResponseWriter, r *http.Request, c middle.Context)
	Update(w http.ResponseWriter, r *http.Request, c middle.Context)
	Delete(w http.ResponseWriter, r *http.Request, c middle.Context)
}

// CreateAllRoutes is a convenience method for subscribing to all routes.
func CreateAllRoutes(model string, handler Handler) []control.Route {
	crud := []Route{
		Route{Op: Create, HandlerFunc: handler.Create},
		Route{Op: Find, HandlerFunc: handler.Find},
		Route{Op: FindOne, HandlerFunc: handler.FindOne},
		Route{Op: Update, HandlerFunc: handler.Update},
		Route{Op: Delete, HandlerFunc: handler.Delete},
	}
	return CreateRoutes(model, crud)
}

// CreateRoutes transforms CRUDROutes to Routes expected by the server.
func CreateRoutes(model string, crud []Route) []control.Route {

	m := make(map[string]*control.Route)
	for _, r := range crud {

		getRoute := func(pattern string) *control.Route {
			route := m[pattern]
			if route == nil {
				route = &control.Route{Path: pattern}
				m[pattern] = route
			}
			return route
		}

		switch r.Op {
		case Create:
			route := getRoute(fmt.Sprintf("/%v", model))
			route.Handlers = append(route.Handlers, control.Handler{Method: "post", HandlerFunc: r.HandlerFunc, Middleware: r.Middleware})
		case Find:
			route := getRoute(fmt.Sprintf("/%v", model))
			route.Handlers = append(route.Handlers, control.Handler{Method: "get", HandlerFunc: r.HandlerFunc, Middleware: r.Middleware})
		case FindOne:
			route := getRoute(fmt.Sprintf("/%v/{id:[0-9]+}", model))
			route.Handlers = append(route.Handlers, control.Handler{Method: "get", HandlerFunc: r.HandlerFunc, Middleware: r.Middleware})
		case Update:
			route := getRoute(fmt.Sprintf("/%v/{id:[0-9]+}", model))
			route.Handlers = append(route.Handlers, control.Handler{Method: "patch", HandlerFunc: r.HandlerFunc, Middleware: r.Middleware})
		case Delete:
			route := getRoute(fmt.Sprintf("/%v/{id:[0-9]+}", model))
			route.Handlers = append(route.Handlers, control.Handler{Method: "delete", HandlerFunc: r.HandlerFunc, Middleware: r.Middleware})
		}
	}

	retRoutes := make([]control.Route, 0, len(m))
	for _, value := range m {
		retRoutes = append(retRoutes, *value)
	}

	return retRoutes
}
