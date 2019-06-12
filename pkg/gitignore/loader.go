package gitignore

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/emman27/gitignore/pkg/log"
	"github.com/sirupsen/logrus"
)

// Loader loads gitignore files based on the gitignore type
type Loader interface {
	Load(context.Context, gitignoreType) (string, error)
}

func NewLoader() (Loader, error) {
	return newGithubLoader()
}

func newGithubLoader() (*githubLoader, error) {
	tr := &http.Transport{
		TLSHandshakeTimeout:   3 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}
	logger, err := log.New()
	if err != nil {
		return nil, err
	}
	return &githubLoader{
		client: client,
		logger: logger,
	}, nil
}

type githubLoader struct {
	client *http.Client
	logger logrus.FieldLogger
}

const githubRawURLTemplate = "https://raw.githubusercontent.com/github/gitignore/master/%s.gitignore"

func (l *githubLoader) Load(ctx context.Context, t gitignoreType) (string, error) {
	url := fmt.Sprintf(githubRawURLTemplate, t)
	l.logger.Debugf("Loading gitignore from %s", url)
	resp, err := l.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type gitignoreType string

const (
	Go gitignoreType = "Go"
)
