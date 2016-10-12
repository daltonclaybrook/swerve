package main

import (
	"fmt"
	"github.com/daltonclaybrook/swerve/control"
	"github.com/daltonclaybrook/swerve/crud"
	"github.com/daltonclaybrook/swerve/middle"
	"github.com/gorilla/mux"
	"net/http"
)

// User is the controller for all /user endpoints.
type User struct{}

// Routes describes all endpoints handled by the User Controller.
func (uc *User) Routes() []control.Route {
	// return crud.CreateRoutes("user", []crud.Route{
	// 	crud.Route{Op: crud.Create, HandlerFunc: uc.Create},
	// })

	return crud.CreateAllRoutes("user", uc)
}

/*
Handlers
*/

func (uc *User) Create(w http.ResponseWriter, r *http.Request, c middle.Context) {
	fmt.Fprintln(w, "create")
}

func (uc *User) Find(w http.ResponseWriter, r *http.Request, c middle.Context) {
	fmt.Fprintln(w, "find")
}

func (uc *User) FindOne(w http.ResponseWriter, r *http.Request, c middle.Context) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, "findOne with id: %v\n", id)
}

func (uc *User) Update(w http.ResponseWriter, r *http.Request, c middle.Context) {
	fmt.Fprintln(w, "update")
}

func (uc *User) Delete(w http.ResponseWriter, r *http.Request, c middle.Context) {
	fmt.Fprintln(w, "delete")
}
