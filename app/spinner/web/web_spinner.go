package web

import (
	"context"
	"fmt"
	"github.com/syke99/sfw/pkg/models"
)

type webSpinner struct {
	web *models.Web
}

func NewWebSpinner(web *models.Web) WebSpinner {
	return &webSpinner{web: web}
}

func (w *webSpinner) Cast(ctx context.Context, msg models.Message) error {
	// Hooking into "arachneos.line" will go here
	fmt.Printf("line %d: %s", msg.ID, msg.Text)
	return nil
}
