package server

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/submodule-org/submodule.go"
)

type testSuite struct {
	suite.Suite
	scope submodule.Scope
	srv   Server
}

func TestServer(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (s *testSuite) SetupTest() {
	s.scope = submodule.CreateScope()
	s.srv = ServerMod.ResolveWith(s.scope)
}
