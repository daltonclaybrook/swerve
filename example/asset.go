package main

import (
	"fmt"
	"github.com/daltonclaybrook/swerve/control"
	"github.com/daltonclaybrook/swerve/crud"
	"github.com/daltonclaybrook/swerve/middle"
	"net/http"
)

type Asset struct{}

func (a *Asset) Routes() []control.Route {
	return crud.CreateRoutes("asset", []crud.Route{
		crud.Route{Op: crud.Find, HandlerFunc: a.Find},
		crud.Route{Op: crud.Create, HandlerFunc: a.Create},
	})
}

func (a *Asset) Find(w http.ResponseWriter, r *http.Request, c middle.Context) {
	fmt.Println("find")
}

func (a *Asset) Create(w http.ResponseWriter, r *http.Request, c middle.Context) {
	fmt.Println("create")
}
