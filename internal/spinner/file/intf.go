package file

import (
	"context"

	"github.com/syke99/sfw/pkg/models"
)

type FileSpinner interface {
	Cast(ctx context.Context, msg models.Message, errs chan<- error)
}
