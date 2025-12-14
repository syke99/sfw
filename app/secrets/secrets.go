package secrets

type secrets struct{}

func NewSecretsStore() SecretsStore {
	// inject dependencies here
	return &secrets{}
}

func (s *secrets) Get(key string) (string, error) {
	return "", nil
}
