package gitignore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestNewParser(t *testing.T) {
	p, err := NewParser()
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

type ParserSuite struct {
	suite.Suite
	parser Parser
}

func TestParserSuite(t *testing.T) {
	suite.Run(t, new(ParserSuite))
}

func (s *ParserSuite) SetupTest() {
	p, err := NewParser()
	s.Suite.Nil(err)
	s.parser = p
}

func (s *ParserSuite) TestGo() {
	s.testHelper([]string{"go", "Go", "golang"}, Go)
}

func (s *ParserSuite) TestPython() {
	s.testHelper([]string{"py", "python", "Python"}, Python)
}

func (s *ParserSuite) TestUnparsable() {
	_, err := s.parser.Parse(context.TODO(), "dsfjaklfjladsfjalj")
	s.NotNil(err)
}

func (s *ParserSuite) testHelper(expectedCanParse []string, shouldParseTo gitignoreType) {
	for _, parsable := range expectedCanParse {
		parsed, err := s.parser.Parse(context.TODO(), parsable)
		s.Nil(err)
		s.Equal(shouldParseTo, parsed)
	}
}
