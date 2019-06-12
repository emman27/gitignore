package gitignore

import (
	"context"
	"errors"
)

// Service provides gitignore generation services
type Service interface {
	Generate(ctx context.Context, gitignoreType string) (string, error)
}

func NewService(options ...ServiceOption) (Service, error) {
	parser, err := NewParser()
	if err != nil {
		return nil, err
	}
	loader, err := NewLoader()
	if err != nil {
		return nil, err
	}
	svc := &service{parser: parser, loader: loader}
	for _, opt := range options {
		if err := opt(svc); err != nil {
			return svc, err
		}
	}
	return svc, nil
}

type ServiceOption func(*service) error

func WithParser(parser Parser) ServiceOption {
	return func(s *service) error {
		if parser == nil {
			return errors.New("invalid parser")
		}
		s.parser = parser
		return nil
	}
}

func WithLoader(loader Loader) ServiceOption {
	return func(s *service) error {
		if loader == nil {
			return errors.New("invalid loader")
		}
		s.loader = loader
		return nil
	}
}

type service struct {
	parser Parser
	loader Loader
}

func (s *service) Generate(ctx context.Context, gitignoreType string) (string, error) {
	ignoreType, err := s.parser.Parse(ctx, gitignoreType)
	if err != nil {
		return "", err
	}
	return s.loader.Load(ctx, ignoreType)
}
