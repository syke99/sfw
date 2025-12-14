package spinner

import (
	"context"
	f "github.com/syke99/sfw/internal/spinner/file"
	ws "github.com/syke99/sfw/internal/spinner/webhook"
	"github.com/syke99/sfw/pkg/models"
)

type Type int

const (
	Web Type = iota
	File
)

type spinnerI interface {
	Cast(ctx context.Context, msg models.Message, errs chan<- error)
}

type spinner struct {
	web    *models.StickyWeb
	state  map[string]string
	s      spinnerI
	st     Type
	source string

	// anything else needed
}

func NewWebSpinner(web *models.Web, stickyWeb *models.StickyWeb, source string) (Spinner, error) {
	return &spinner{
		web:    stickyWeb,
		state:  make(map[string]string),
		s:      ws.NewWebhookSpinner(web),
		st:     Web,
		source: source,
	}, nil
}

func NewFileSpinner(web *models.Web, stickyWeb *models.StickyWeb, source string) (Spinner, error) {
	return &spinner{
		web:    stickyWeb,
		state:  make(map[string]string),
		s:      f.NewFileSpinner(web),
		st:     File,
		source: source,
	}, nil
}

func (s *spinner) Cast(ctx context.Context, msg models.Message, errs chan<- error) {
	s.s.Cast(ctx, msg, errs)
}

func (s *spinner) Type() Type {
	return s.st
}

func (s *spinner) Source() string {
	return s.source
}
