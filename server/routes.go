package server

import (
	"reference/common"
	"reference/todo_route"

	"github.com/submodule-org/submodule.go"
)

var routes = submodule.Group[[]common.Registry](todo_route.Registries)
