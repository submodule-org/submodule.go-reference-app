package todo_route

import (
	"fmt"
	"net/http"

	"reference/common"
	"reference/todo_svc"

	"github.com/submodule-org/submodule.go"
)

var get = submodule.Make[common.Registry](func(svc todo_svc.TodoSvc) common.Registry {
	return common.Registry{
		Path: "GET /todo/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "GET Hello World!")
		},
	}
}, todo_svc.TodoSvcMod)

var insert = submodule.Make[common.Registry](func() common.Registry {
	return common.Registry{
		Path: "POST /todo",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "POST Hello World!")
		},
	}
}, todo_svc.TodoSvcMod)

var delete = submodule.Make[common.Registry](func() common.Registry {
	return common.Registry{
		Path: "DELETE /todo/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "DELETE Hello World!")
		},
	}
}, todo_svc.TodoSvcMod)

var Route = common.MakeMuxMod("/todo", get, insert, delete)
