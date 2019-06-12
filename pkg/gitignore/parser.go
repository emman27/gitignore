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

func NewParser() (Parser, error) {
	return &parser{}, nil
}

type parser struct{}

func (p *parser) Parse(ctx context.Context, s string) (gitignoreType, error) {
	return gitignoreType(strings.Title(s)), nil
}
