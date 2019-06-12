package gitignore

import (
	"context"
	"fmt"
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

var parserMap = map[string]gitignoreType{
	"go":     Go,
	"Go":     Go,
	"golang": Go,
}

func (p *parser) Parse(ctx context.Context, s string) (gitignoreType, error) {
	parsed, ok := parserMap[s]
	if !ok {
		return "", fmt.Errorf("%s is not a recognized gitignore type", s)
	}
	return parsed, nil
}
