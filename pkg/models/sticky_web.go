package models

type StickyWeb struct {
	Web     *Web
	Secrets map[string]string
	Inputs  map[string][]string
	Outputs map[string][]string
}
