package log

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var logLevel = flag.String("log-level", "info", "Logging level to use")

func New() (*logrus.Logger, error) {
	logger := logrus.New()
	level, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		logger.Errorf("Log level not accepted, falling back to info. You provided: %s", *logLevel)
		return logger, nil
	}
	logger.SetLevel(level)
	return logger, nil
}
