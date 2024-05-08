package todo_route

import (
	"fmt"
	"net/http"

	rh "reference/http"
	"reference/todo_svc"

	"github.com/submodule-org/submodule.go"
)

var get = submodule.Make[rh.Registry](func(svc todo_svc.TodoSvc) rh.Registry {
	return rh.Registry{
		Path: "GET /todo/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "GET Hello World!")
		},
	}
}, todo_svc.TodoSvcMod)

var insert = submodule.Make[rh.Registry](func() rh.Registry {
	return rh.Registry{
		Path: "POST /todo",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "POST Hello World!")
		},
	}
}, todo_svc.TodoSvcMod)

var delete = submodule.Make[rh.Registry](func() rh.Registry {
	return rh.Registry{
		Path: "DELETE /todo/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "DELETE Hello World!")
		},
	}
}, todo_svc.TodoSvcMod)

var Route = rh.MakeMuxMod("/todo", get, insert, delete)
