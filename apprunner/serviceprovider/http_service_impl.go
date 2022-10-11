package serviceprovider

import (
	"context"
	"fmt"
	"net/http"
)

type HttpService struct {
	Http *http.Server
}

func (g HttpService) Name() string {
	return "http"
}

func (g HttpService) Start() error {
	err := g.Http.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("http: start server: %w", err)
	}
	return nil
}

func (g HttpService) Stop(ctx context.Context) error {
	return g.Http.Shutdown(ctx)
}

func (g HttpService) BeforeStop() {}
