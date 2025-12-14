package models

import (
	"github.com/syke99/sfw/internal/pkg/models"
)

type Web struct {
	Web   *models.Web
	Lines []*models.Line
	Knots map[string][]*models.Knot
}
