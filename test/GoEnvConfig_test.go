package test

import (
	"github.com/j7mbo/goenvconfig"
	"github.com/j7mbo/goenvconfig/test/deps"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type InjectorTestSuite struct {
	suite.Suite
	parser goenvconfig.GoEnvParser
}

func TestInjectorTestSuite(t *testing.T) {
	tests := new(InjectorTestSuite)

	suite.Run(t, tests)
}

func (s *InjectorTestSuite) SetupSuite() {
	s.parser = goenvconfig.NewGoEnvParser()
}

func (s *InjectorTestSuite) TestCanSetStringVarFromEnv() {
	dep := deps.Dep{}

	_ = os.Setenv("s", "TEST")

	_ = s.parser.Parse(&dep)

	s.Assert().Equal(dep.GetString(), "TEST")
}

func (s *InjectorTestSuite) TestCanSetIntVarFromEnv() {
	dep := deps.Dep{}

	_ = os.Setenv("i", "42")

	_ = s.parser.Parse(&dep)

	s.Assert().Equal(dep.GetInt(), 42)
}

func (s *InjectorTestSuite) TestDefaultUsedWhenNoEnvVarFound() {
	dep := deps.Dep{}

	_ = s.parser.Parse(&dep)

	s.Assert().Equal(dep.GetInt(), 1337)
}

func (s *InjectorTestSuite) TestPassingInAStructValueReturnsError() {
	dep := deps.Dep{}

	err := s.parser.Parse(dep)

	s.Assert().Error(err)
}

func (s *InjectorTestSuite) TestCanStillDoPublicPropertiesAlso() {
	dep := deps.Dep{}

	_ = os.Setenv("ip", "1338")

	_ = s.parser.Parse(&dep)

	s.Assert().Equal(dep.GetPublicInt(), 1338)
}

func (s *InjectorTestSuite) TearDownTest() {
	_ = os.Unsetenv("t")
	_ = os.Unsetenv("i")
	_ = os.Unsetenv("ip")
}
