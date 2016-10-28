package main

import (
	"fmt"
	"github.com/daltonclaybrook/swerve/middle"
	"net/http"
)

type IsAuthenticated struct{}
type IsOwnerOfAsset struct{}

func (ia IsAuthenticated) Handle(w http.ResponseWriter, r *http.Request, context middle.Context, next middle.NextFunc) {
	fmt.Println("is authenticated policy...")
	next(context)
}

func (iooa IsOwnerOfAsset) Handle(w http.ResponseWriter, r *http.Request, context middle.Context, next middle.NextFunc) {
	fmt.Println("is owner of asset policy...")
	next(context)
}
