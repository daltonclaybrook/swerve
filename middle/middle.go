package middle

import (
	"net/http"
)

// Handler is an interface which middleware must conform to
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request, context Context, next NextFunc)
}

// Context is a map used to carry state from middleware to a control
type Context map[string]interface{}

// ContextFunc is effectively an http.HandlerFunc plus a context
type ContextFunc func(http.ResponseWriter, *http.Request, Context)

// NextFunc is a function passed to middleware to call when the middleware operation is finished
type NextFunc func(Context)

// CreateHandlerFunc is used to chain an array of Handlers in front of a terminating ContextFunc
func CreateHandlerFunc(handlers []Handler, final ContextFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next := nextHandlerFunc(handlers, final, 0, w, r)
		next(make(Context))
	}
}

/*
Private
*/

func nextHandlerFunc(handlers []Handler, final ContextFunc, handlerIdx int, w http.ResponseWriter, r *http.Request) NextFunc {
	return func(c Context) {
		if handlerIdx < len(handlers) {
			handler := handlers[handlerIdx]
			next := nextHandlerFunc(handlers, final, handlerIdx+1, w, r)
			handler.Handle(w, r, c, next)
		} else {
			final(w, r, c)
		}
	}
}
