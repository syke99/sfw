package web

import (
	"context"
	"fmt"
	"github.com/syke99/sfw/app/spinner"

	"github.com/go-chi/chi/v5"

	"github.com/syke99/sfw/app/web/internal"
	"github.com/syke99/sfw/internal/web"
)

type wb struct {
	web web.WebCaster
	// this will start up the individual
	// goroutines to handle each spinner,
	// put configuration here
}

// is core of the application; injecting
// spinners is how people will be able to
// hook into app with thier own implementations
func NewWeb(mux *chi.Mux, path string) (WebCaster, error) {
	spinners, err := internal.BuildSpinners(path)
	if err != nil {
		err = fmt.Errorf("error building spinners: %w", err)
		return nil, err
	}

	return &wb{
		web: web.NewWeb(mux, spinners),
	}, nil
}

func NewWebWithSpinners(mux *chi.Mux, spinners map[string]spinner.Spinner) (WebCaster, error) {
	return &wb{
		web: web.NewWeb(mux, spinners),
	}, nil
}

func (s *wb) Cast(ctx context.Context) error {
	return s.web.Cast(ctx)
}
