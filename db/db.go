package db

import (
	"github.com/hashicorp/go-memdb"
	"github.com/submodule-org/submodule.go"
)

var Db = func(schema *memdb.DBSchema) submodule.Submodule[*memdb.MemDB] {
	return submodule.Make[*memdb.MemDB](func() (*memdb.MemDB, error) {
		return memdb.NewMemDB(schema)
	})
}
