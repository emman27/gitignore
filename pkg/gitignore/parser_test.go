package gitignore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	p, err := NewParser()
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func Test_parser_Parse(t *testing.T) {
	p := &parser{}
	parsed, err := p.Parse(context.TODO(), "go")
	assert.Nil(t, err)
	assert.Equal(t, Go, parsed)
}
