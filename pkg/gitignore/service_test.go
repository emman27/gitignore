package gitignore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

//go:generate mockery -name=Loader -inpkg -testonly

type ServiceSuite struct {
	suite.Suite
	service Service
	loader  *MockLoader
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	s.loader = &MockLoader{}
	svc, err := NewService(WithLoader(s.loader))
	s.Nil(err)
	s.service = svc
}

func (s *ServiceSuite) TestGenerate() {
	ctx := context.TODO()
	fakeGitignore := "fake gitignore"
	s.loader.On("Load", ctx, Go).Return(fakeGitignore, nil)
	result, err := s.service.Generate(context.TODO(), "go")
	s.Nil(err)
	s.Equal(fakeGitignore, result)
}

func (s *ServiceSuite) TestParseFailure() {
	_, err := s.service.Generate(context.TODO(), "dklajflsfj")
	s.NotNil(err)
}

func (s *ServiceSuite) TestNilLoader() {
	_, err := NewService(WithLoader(nil))
	s.NotNil(err)
}

func (s *ServiceSuite) TestNilParser() {
	_, err := NewService(WithParser(nil))
	s.NotNil(err)
}

func (s *ServiceSuite) TestWithParser() {
	p, err := NewParser()
	s.Nil(err)
	_, err = NewService(WithParser(p))
	s.Nil(err)
}
