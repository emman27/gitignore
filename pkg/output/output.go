package output

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/emman27/gitignore/pkg/log"
)

const (
	outputTypeGitIgnoreFile = "gitignore"
	outputTypeConsole       = "console"
)

var outputType = flag.String("output-type", outputTypeGitIgnoreFile, "Output type. Can be gitignore or console")

type Output interface {
	Write(context.Context, string) error
}

func New() (Output, error) {
	switch *outputType {
	case outputTypeGitIgnoreFile:
		return &gitIgnoreFile{}, nil
	case outputTypeConsole:
		return &console{}, nil
	default:
		logger, err := log.New()
		if err != nil {
			return nil, err
		}
		logger.Warnf("Unknown output type %s, defaulting to console", *outputType)
		return &console{}, nil
	}
}

type gitIgnoreFile struct{}

func (f *gitIgnoreFile) Write(ctx context.Context, content string) error {
	return ioutil.WriteFile(".gitignore", []byte(content), 0644)
}

type console struct{}

func (c *console) Write(ctx context.Context, content string) error {
	fmt.Println(content)
	return nil
}
