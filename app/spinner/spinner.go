package spinner

import (
	"context"
	"fmt"

	p "github.com/syke99/sfw/app/parser"
	f "github.com/syke99/sfw/app/spinner/file"
	ws "github.com/syke99/sfw/app/spinner/webhook"
	"github.com/syke99/sfw/pkg/models"
)

type spinnerI interface {
	Cast(ctx context.Context, msg models.Message, errs chan<- error)
}

type spinner struct {
	web   *p.StickyWeb
	state map[string]string
	s     spinnerI
	st    Type

	// anything else needed
}

func NewWebSpinner(web *models.Web, parser p.Parser) (Spinner, error) {
	stickyWeb, err := parser.Parse(web)
	if err != nil {
		err = fmt.Errorf("failed to make web sticky: %w", err)
		return nil, err
	}

	return &spinner{
		web:   stickyWeb,
		state: make(map[string]string),
		s:     ws.NewWebhookSpinner(web),
		st:    Web,
	}, nil
}

func NewFileSpinner(web *models.Web, parser p.Parser) (Spinner, error) {
	stickyWeb, err := parser.Parse(web)
	if err != nil {
		// TODO: wrap error
		return nil, err
	}

	return &spinner{
		web:   stickyWeb,
		state: make(map[string]string),
		s:     f.NewFileSpinner(web),
		st:    File,
	}, nil
}

func (s *spinner) Cast(ctx context.Context, msg models.Message, errs chan<- error) {
	s.s.Cast(ctx, msg, errs)
}

type Type int

const (
	Web Type = iota
	File
)

func (s *spinner) Type() Type {
	return s.st
}
