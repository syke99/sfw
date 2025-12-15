package models

import (
	"github.com/syke99/sfw/internal/pkg/models"
)

type Web struct {
	Web   *models.Web
	Lines []*models.Line
	// groups of knots by line
	Knots map[string][]*models.Knot
}
