package server

import (
	"reference/http"
	"reference/todo_route"

	"github.com/submodule-org/submodule.go"
)

var routes = submodule.Group[http.Mux](todo_route.Route)
