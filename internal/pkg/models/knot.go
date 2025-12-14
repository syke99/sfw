package models

type Knot struct {
	Name          string `yaml:"name"`
	Knot          string `yaml:"knot"`
	Observability struct {
		Provider interface{} `yaml:"provider"`
	} `yaml:"observability"`
	MemoryLimit  int               `yaml:"memoryLimit"`
	AllowedUrls  []string          `yaml:"allowedUrls"`
	AllowedPaths map[string]string `yaml:"allowedPaths"`
	Secrets      []string          `yaml:"secrets"`
	Inputs       []string          `yaml:"inputs"`
	Outputs      []string          `yaml:"outputs"`
}
