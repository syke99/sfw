package parser

import "github.com/syke99/sfw/pkg/models"

type Parser interface {
	Parse(web *models.Web) (*StickyWeb, error)
}
