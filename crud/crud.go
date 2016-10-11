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
	Op      int
	Handler middle.ContextFunc
}

// Handler defines a type which handles all CRUD methods.
type Handler interface {
	create(w http.ResponseWriter, r *http.Request, c middle.Context)
	find(w http.ResponseWriter, r *http.Request, c middle.Context)
	findOne(w http.ResponseWriter, r *http.Request, c middle.Context)
	update(w http.ResponseWriter, r *http.Request, c middle.Context)
	delete(w http.ResponseWriter, r *http.Request, c middle.Context)
}

// AllRoutesFromHandler is a convenience method for subscribing to all routes.
func AllRoutesFromHandler(model string, handler Handler) []control.Route {
	crud := []Route{
		Route{Create, handler.create},
		Route{Find, handler.find},
		Route{FindOne, handler.findOne},
		Route{Update, handler.update},
		Route{Delete, handler.delete},
	}
	return RoutesFromCRUD(model, crud)
}

// RoutesFromCRUD transforms CRUDROutes to Routes expected by the server.
func RoutesFromCRUD(model string, crud []Route) []control.Route {

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
			route.Handlers = append(route.Handlers, control.Handler{Method: "post", Handler: r.Handler})
		case Find:
			route := getRoute(fmt.Sprintf("/%v", model))
			route.Handlers = append(route.Handlers, control.Handler{Method: "get", Handler: r.Handler})
		case FindOne:
			route := getRoute(fmt.Sprintf("/%v/{id:[0-9]+}", model))
			route.Handlers = append(route.Handlers, control.Handler{Method: "get", Handler: r.Handler})
		case Update:
			route := getRoute(fmt.Sprintf("/%v/{id:[0-9]+}", model))
			route.Handlers = append(route.Handlers, control.Handler{Method: "patch", Handler: r.Handler})
		case Delete:
			route := getRoute(fmt.Sprintf("/%v/{id:[0-9]+}", model))
			route.Handlers = append(route.Handlers, control.Handler{Method: "delete", Handler: r.Handler})
		}
	}

	retRoutes := make([]control.Route, 0, len(m))
	for _, value := range m {
		retRoutes = append(retRoutes, *value)
	}

	return retRoutes
}
