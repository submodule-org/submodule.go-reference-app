package todo_svc

import (
	"context"
	"reference/common"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/submodule-org/submodule.go"
)

type testSuite struct {
	suite.Suite
	scope submodule.Scope
	svc   TodoSvc
}

func (s *testSuite) SetupTest() {
	s.scope = submodule.CreateScope()
	s.svc = TodoSvcMod.ResolveWith(s.scope)
}

func (s *testSuite) TestCRUD() {
	var todo common.Todo
	var err error

	todo, err = s.svc.Insert(context.Background(), common.Todo{
		Content: "example",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), todo.Id)

	todo, err = s.svc.Get(context.Background(), todo.Id)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "example", todo.Content)

	todo, err = s.svc.Delete(context.Background(), todo.Id)
	assert.Nil(s.T(), err)

	todo, err = s.svc.Get(context.Background(), todo.Id)
	assert.NotNil(s.T(), err)
}

func TestTodoSvc(t *testing.T) {
	suite.Run(t, new(testSuite))
}
