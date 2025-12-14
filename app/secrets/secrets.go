package secrets

import "github.com/syke99/sfw/pkg/models"

type secrets struct{}

func NewSecretsStore(web *models.Web) SecretsStore {
	// inject dependencies here
	return &secrets{}
}

func (s *secrets) Get(key string) (string, error) {
	return "", nil
}
