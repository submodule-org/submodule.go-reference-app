package todo_route

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"reference/common"
	"reference/http_test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/submodule-org/submodule.go"
)

type testSuite struct {
	suite.Suite
	scope  submodule.Scope
	port   int
	srv    http.Server
	client http_test.Client
}

func (s *testSuite) SetupTest() {
	s.port = common.NextInt(40000, 50000)

	s.scope = submodule.CreateScope()
	route := Route.ResolveWith(s.scope)

	r := &http.ServeMux{}
	r.Handle(route.Path, route.Mux)

	s.srv = http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: r,
	}

	socket, e := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if e != nil {
		panic(e)
	}

	go func() {
		_ = s.srv.Serve(socket)
	}()

	s.client = http_test.NewClient(http_test.WithPort(s.port))
}

func (s *testSuite) TearDownTest() {
	s.srv.Shutdown(context.TODO())
}

func (s *testSuite) TestGet() {
	response, err := s.client.Get(fmt.Sprintf("/todo/%s", common.NextId()))
	assert.Nil(s.T(), err)

	responseData, err := io.ReadAll(response.Body)
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), `GET Hello World!`, string(responseData))
}

func TestTodoRoute(t *testing.T) {
	suite.Run(t, new(testSuite))
}
