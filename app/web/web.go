package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/syke99/sfw/app/spinner"
	"github.com/syke99/sfw/app/web/internal"
	"github.com/syke99/sfw/internal/web"
)

type wb struct {
	web web.WebCaster
	// this will start up the individual
	// goroutines to handle each spinner,
	// put configuration here
}

func NewWeb(mux http.Handler, path string) (WebCaster, error) {
	spinners, err := internal.BuildSpinners(path)
	if err != nil {
		err = fmt.Errorf("error building spinners: %w", err)
		return nil, err
	}

	return &wb{
		web: web.NewWeb(mux, spinners),
	}, nil
}

// NewWebWithSpinners is core of the application;
// injecting your own spinner.Spinner interface is
// how you can create custom hooks into the WebCaster
// if embedding <app-name-here> your own system,
// use the engine to build from scratch, or
// anything else
func NewWebWithSpinners(mux http.Handler, spinners []spinner.Spinner) (WebCaster, error) {
	return &wb{
		web: web.NewWeb(mux, spinners),
	}, nil
}

func (s *wb) Cast(ctx context.Context) error {
	return s.web.Cast(ctx)
}
