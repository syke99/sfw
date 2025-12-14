package spinner

import (
	"context"
	"github.com/syke99/sfw/pkg/models"
)

type Spinner interface {
	Cast(ctx context.Context, msg models.Message, errs chan<- error)
	Type() Type
	Source() string
}
