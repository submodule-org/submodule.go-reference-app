package todo_svc

import (
	"context"
	"reference/common"

	"github.com/hashicorp/go-memdb"
	"github.com/submodule-org/submodule.go"
)

type todoSvc struct {
	db *memdb.MemDB
}

type TodoSvc interface {
	Insert(ctx context.Context, todo common.Todo) (common.Todo, error)
	Get(ctx context.Context, id string) (common.Todo, error)
	Delete(ctx context.Context, id string) (common.Todo, error)
}

func (t *todoSvc) Insert(ctx context.Context, todo common.Todo) (c common.Todo, e error) {
	txn := t.db.Txn(true)

	nextId := nextId()
	todo.Id = nextId

	e = txn.Insert("todo", todo)

	if e != nil {
		return
	}

	c = todo
	return
}

func (t *todoSvc) Get(ctx context.Context, id string) (c common.Todo, e error) {
	txn := t.db.Txn(false)
	defer txn.Abort()

	x, e := txn.First("todo", id)
	return x.(common.Todo), nil
}

func (t *todoSvc) Delete(ctx context.Context, id string) (c common.Todo, e error) {
	txn := t.db.Txn(true)
	defer txn.Abort()

	x, e := txn.First("todo", id)
	if e != nil {
		return
	}

	todo := x.(common.Todo)
	todo.Done = true

	e = txn.Delete("todo", todo)
	if e != nil {
		return
	}

	return todo, nil
}

var TodoSvcMod = submodule.Resolve[TodoSvc](&todoSvc{}, todoDb)
