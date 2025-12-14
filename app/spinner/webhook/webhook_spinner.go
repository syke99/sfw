package webhook

import (
	"context"
	"fmt"
	"github.com/syke99/sfw/pkg/models"
)

type webhookSpinner struct {
	web *models.Web
}

func NewWebhookSpinner(web *models.Web) WebhookSpinner {
	return &webhookSpinner{web: web}
}

func (w *webhookSpinner) Cast(ctx context.Context, msg models.Message) error {
	// Hooking into "arachneos.line" will go here
	fmt.Printf("line %d: %s", msg.ID, msg.Text)
	return nil
}
