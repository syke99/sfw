package web

import (
	"context"

	"github.com/syke99/sfw/pkg/models"
)

type WebSpinner interface {
	Cast(ctx context.Context, msg models.Message) error
}
