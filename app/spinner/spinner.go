package spinner

import (
	"context"

	p "github.com/syke99/sfw/app/parser"
	f "github.com/syke99/sfw/app/spinner/file"
	ws "github.com/syke99/sfw/app/spinner/web"
	"github.com/syke99/sfw/pkg/models"
)

type spinner struct {
	web    *models.Web
	s      Spinner
	parser p.Parser

	// anything else needed
}

func NewWebSpinner(web *models.Web, parser p.Parser) Spinner {
	return &spinner{
		web:    web,
		s:      ws.NewWebSpinner(web),
		parser: parser,
	}
}

func NewFileSpinner(web *models.Web, parser p.Parser) Spinner {
	return &spinner{
		web:    web,
		s:      f.NewFileSpinner(web),
		parser: parser,
	}
}

func (s *spinner) Cast(ctx context.Context, msg models.Message) error {
	return s.s.Cast(ctx, msg)
}
