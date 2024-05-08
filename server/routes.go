package server

import (
	"reference/common"
	"reference/todo_route"

	"github.com/submodule-org/submodule.go"
)

var routes = submodule.Group[common.Mux](todo_route.Route)
