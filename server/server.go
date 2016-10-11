package server

import (
	"fmt"
	"github.com/daltonclaybrook/swerve/control"
	"github.com/daltonclaybrook/swerve/middle"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// WebServer is used to create and start a server.
type WebServer struct {
	server     *http.Server
	router     *mux.Router
	middleware []middle.Handler
	controls   []control.Control
}

// NewWebServer returns a new initialized instance of WebServer.
func NewWebServer() *WebServer {
	ws := &WebServer{}
	ws.controls = make([]control.Control, 0)
	ws.middleware = make([]middle.Handler, 0)
	ws.router = mux.NewRouter()

	http.Handle("/", ws.router)
	return ws
}

// RegisterControl registers a request handler with the WebServer.
func (ws *WebServer) RegisterControl(c control.Control) {
	ws.controls = append(ws.controls, c)
}

// RegisterMiddleware registers request handlers called before the control.
func (ws *WebServer) RegisterMiddleware(m middle.Handler) {
	ws.middleware = append(ws.middleware, m)
}

// Start starts the server.
func (ws *WebServer) Start() {
	ws.setupServer()
	ws.addRoutesForControls()
	ws.server.ListenAndServe()
}

/*
Private
*/

func (ws *WebServer) setupServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ws.server = &http.Server{
		Addr:           fmt.Sprintf(":%v", port),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	ws.registerHandler("/", sendUnhandled, "")
	ws.server.ErrorLog = log.New(os.Stdout, "err: ", 0)
	// ws.server.ConnState = func(con net.Conn, state http.ConnState) {
	// 	fmt.Printf("con: %v, state: %v\n", con, state)
	// }
}

func (ws *WebServer) addRoutesForControls() {
	for _, c := range ws.controls {
		routes := c.Routes()
		for _, route := range routes {
			ws.registerRouteHandlers(route)
		}
	}
}

func (ws *WebServer) registerRouteHandlers(route control.Route) {
	methods := make([]string, len(route.Handlers))
	for idx, handler := range route.Handlers {
		fmt.Printf("path: %v, method: %v\n", route.Path, handler.Method)

		ws.registerHandler(route.Path, handler.Handler, handler.Method)
		methods[idx] = strings.ToUpper(handler.Method)
	}
	ws.registerHandler(route.Path, sendOptionsHandlerFunc(methods), "options")
}

func (ws *WebServer) registerHandler(path string, handlerFunc middle.ContextFunc, method string) {
	toRegister := middle.HandlerFuncByApplyingMiddleware(ws.middleware, handlerFunc)
	route := ws.router.HandleFunc(path, toRegister)
	if len(method) > 0 {
		route.Methods(method)
	}
}

func sendUnhandled(w http.ResponseWriter, r *http.Request, c middle.Context) {
	w.WriteHeader(404)
	fmt.Fprintf(w, "Method \"%v\" is not supported by this route.", r.Method)
}

func sendOptionsHandlerFunc(methods []string) func(w http.ResponseWriter, r *http.Request, c middle.Context) {
	methods = append(methods, "OPTIONS")
	methodString := strings.Join(methods, ", ")
	return func(w http.ResponseWriter, r *http.Request, c middle.Context) {
		w.Header().Set("Allow", methodString)
		w.Header().Set("Content-Length", "0")
	}
}