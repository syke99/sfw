package file

import (
	"context"
	"fmt"
	"github.com/syke99/sfw/pkg/models"
)

type fileSpinner struct {
	web *models.Web
}

func NewFileSpinner(web *models.Web) FileSpinner {
	return &fileSpinner{web: web}
}

func (w *fileSpinner) Cast(ctx context.Context, msg models.Message) error {
	// Hooking into "arachneos.line" will go here
	fmt.Printf("line %d: %s", msg.ID, msg.Text)
	return nil
}
