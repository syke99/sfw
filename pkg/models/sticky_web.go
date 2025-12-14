package models

type StickyWeb struct {
	web     *Web
	secrets map[string]string
	inputs  map[string][]string
	outputs map[string][]string
}
