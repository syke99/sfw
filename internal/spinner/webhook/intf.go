package webhook

import (
	"context"

	"github.com/syke99/sfw/pkg/models"
)

type WebhookSpinner interface {
	Cast(ctx context.Context, msg models.Message, errs chan<- error)
}
