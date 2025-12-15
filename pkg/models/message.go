package models

import extism "github.com/extism/go-sdk"

type Message struct {
	ID   int
	Data []byte
}

type MessageManifest struct {
	WasmPath      string
	Observability string
	MemoryLimit   int
	AllowedUrls   []string
	AllowedPaths  map[string]string
	HostFuncs     []extism.HostFunction
	TraceMetaData map[string]string
}
