package models

type StickyWeb struct {
	Web       *Web
	Secrets   map[string]string
	Inputs    map[string][]string
	Outputs   map[string][]string
	Manifests map[string]*MessageManifest
	States    map[string]State
}

type State = map[string][]byte
