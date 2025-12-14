package spinner

import (
	"context"
	"net/http"

	"github.com/syke99/sfw/pkg/models"
)

type Spinner interface {
	Cast(ctx context.Context, msg models.Message, errs chan<- error)
	Type() Type
	Source() string
}

type Router interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}
