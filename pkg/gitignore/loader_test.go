package gitignore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoaderSuite struct {
	suite.Suite
	loader Loader
}

func TestLoaderSuite(t *testing.T) {
	suite.Run(t, new(LoaderSuite))
}

func (s *LoaderSuite) SetupTest() {
	loader, err := NewLoader()
	s.Nil(err)
	s.loader = loader
}

func (s *LoaderSuite) TestLoadGo() {
	if testing.Short() {
		s.T().Skip()
	}
	result, err := s.loader.Load(context.TODO(), Go)
	s.Nil(err)
	s.NotEmpty(result)
	s.Contains(result, "*.exe", "A golang gitignore file should definitely include the .exe pattern, since go build generates exe files on windows")
}

// Load a gitignore that doesn't exist. Abuses the encapsulation by passing a string instead of a gitignoreType here
func (s *LoaderSuite) TestLoadRubbish() {
	if testing.Short() {
		s.T().Skip()
	}
	_, err := s.loader.Load(context.TODO(), "ajklfjkjva")
	s.NotNil(err)
}
