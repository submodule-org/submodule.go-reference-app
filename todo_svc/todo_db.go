package todo_svc

import (
	"reference/db"

	"github.com/hashicorp/go-memdb"
)

var todoSchema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"todo": {
			Name: "todo",
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "Id"},
				},
			},
		},
	},
}

var todoDb = db.Db(todoSchema)
