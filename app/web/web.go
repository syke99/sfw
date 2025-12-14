package web

import (
	"context"

	"github.com/go-chi/chi/v5"

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
func NewWebFromSpinners(mux *chi.Mux, path string) WebCaster {
	// TODO: open path, parse internal
	// TODO: path, build spinners from
	// TODO: there, and then build
	// TODO: internal web

	return &wb{
		web: web.NewWeb(mux, nil),
	}
}

func (s *wb) Cast(ctx context.Context) error {
	return s.web.Cast(ctx)
}
