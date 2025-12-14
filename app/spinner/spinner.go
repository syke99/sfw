package spinner

import (
	"context"

	ws "github.com/syke99/sfw/app/spinner/web"
	"github.com/syke99/sfw/pkg/models"
)

type spinner struct {
	web *models.Web
	// anything else needed
	s Spinner
}

func NewWebSpinner(web *models.Web) Spinner {
	return &spinner{
		web: web,
		s:   ws.NewWebSpinner(web),
	}
}

func (s *spinner) Cast(ctx context.Context, msg models.Message) error {
	return s.s.Cast(ctx, msg)
}
