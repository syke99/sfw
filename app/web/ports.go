package web

import "context"

type WebCaster interface {
	Cast(ctx context.Context) error
}
