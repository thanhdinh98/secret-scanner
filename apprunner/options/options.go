package options

import (
	"os"
	"time"

	"guardian/config"
)

type Options struct {
	GracefulTimeout time.Duration
}

func Parse() Options {
	gracefulTimeout, _ := time.ParseDuration(os.Getenv(config.KeyGracefulTimeout))
	return Options{
		GracefulTimeout: gracefulTimeout,
	}
}
