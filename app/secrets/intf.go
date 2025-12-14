package secrets

// TODO: put full secrets retrieval here

type SecretsStore interface {
	Get(key string) (string, error)
}
