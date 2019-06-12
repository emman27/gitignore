package gitignore

import (
	"context"
	"strings"
)

// Parser converts strings to their corresponding gitignore type
// Allows mapping of multiple strings to a single, well known gitignore type
// For example,
//   parser(ctx, "golang")
//   "Go"
//   parser(ctx, "go")
//   "Go"
type Parser interface {
	Parse(ctx context.Context, s string) (gitignoreType, error)
}

// NewParser initializes a new parser
func NewParser() (Parser, error) {
	return &parser{}, nil
}

// Basic parser, assumes that the gitignoreType is simply a title casing of the input string
type parser struct{}

func (p *parser) Parse(ctx context.Context, s string) (gitignoreType, error) {
	return gitignoreType(strings.Title(s)), nil
}
