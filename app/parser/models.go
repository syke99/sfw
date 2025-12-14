package parser

import "github.com/syke99/sfw/pkg/models"

type StickyWeb struct {
	web     *models.Web
	secrets map[string]struct{}
	inputs  map[string]struct{}
	outputs map[string]struct{}
}
