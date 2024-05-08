package todo_svc

import (
	"context"
	"errors"
	"reference/common"

	"github.com/hashicorp/go-memdb"
	"github.com/submodule-org/submodule.go"
)

type todoSvc struct {
	Db *memdb.MemDB
}

type TodoSvc interface {
	Insert(ctx context.Context, todo common.Todo) (common.Todo, error)
	Get(ctx context.Context, id string) (common.Todo, error)
	Delete(ctx context.Context, id string) (common.Todo, error)
}

func (t *todoSvc) Insert(ctx context.Context, todo common.Todo) (c common.Todo, e error) {
	txn := t.Db.Txn(true)
	defer txn.Commit()

	nextId := common.NextId()
	todo.Id = nextId

	e = txn.Insert("todo", todo)

	if e != nil {
		return
	}

	c = todo
	return
}

func (t *todoSvc) Get(ctx context.Context, id string) (c common.Todo, e error) {
	txn := t.Db.Txn(false)
	defer txn.Abort()

	x, e := txn.First("todo", "id", id)
	if e != nil {
		return
	}

	if x == nil {
		return c, errors.New("todo not found")
	}

	return x.(common.Todo), nil
}

func (t *todoSvc) Delete(ctx context.Context, id string) (c common.Todo, e error) {
	txn := t.Db.Txn(true)
	defer txn.Commit()

	todo, e := t.Get(ctx, id)
	if e != nil {
		return
	}

	todo.Done = true

	e = txn.Delete("todo", todo)
	if e != nil {
		return
	}

	return todo, nil
}

var TodoSvcMod = submodule.Resolve[TodoSvc](&todoSvc{}, todoDb)
