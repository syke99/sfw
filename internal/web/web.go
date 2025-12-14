package web

import (
	"context"

	"github.com/go-chi/chi/v5"

	"github.com/syke99/sfw/app/spinner"
)

type wb struct {
	mux      *chi.Mux
	spinners map[string]spinner.Spinner
	// this will start up the individual
	// goroutines to handle each spinner,
	// put configuration here
}

// is core of the application; injecting
// spinners is how people will be able to
// hook into app with thier own implementations
func NewWeb(mux *chi.Mux, spinners map[string]spinner.Spinner) WebCaster {
	return &wb{
		mux:      mux,
		spinners: spinners,
	}
}

func (s *wb) Cast(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errs := make(chan error)

	id := 0
	// TODO: NOTE: before starting each spinner, first pre
	// TODO: compile them so we don't have to cold-start
	// TODO: them each time a spinner hooks a models.Message
	for source, sp := range s.spinners {
		lines := make(chan []byte)

		go func() {
			// pass in related plugins and host functions
			s.startSpinner(ctx, id, sp, lines, errs)
		}()

		go func() {
			s.startSpinnerSource(ctx, source, sp.Type(), lines, errs)
		}()

		id++
	}

	select {
	case err := <-errs:
		cancel()
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
