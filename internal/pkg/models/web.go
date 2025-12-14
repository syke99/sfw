package models

type Web struct {
	Version       string `yaml:"version"`
	EngineVersion string `yaml:"engineVersion"`
	Port          int    `yaml:"port"`
	Secrets       struct {
		Provider string `yaml:"provider"`
		Vault    struct {
			URL      string `yaml:"url"`
			Mount    string `yaml:"mount"`
			RoleID   string `yaml:"role_id"`
			SecretID string `yaml:"secret_id"`
		} `yaml:"vault"`
	} `yaml:"secrets"`
	Lines []string `yaml:"lines"`
}
