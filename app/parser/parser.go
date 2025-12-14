package parser

import "github.com/syke99/sfw/app/secrets"

type parser struct {
	secretsStore secrets.SecretsStore
}

func (p parser) ParseSecrets() error {
	//TODO implement parsing here
	return nil
}

func NewParser(secretsStore secrets.SecretsStore) Parser {
	return &parser{secretsStore: secretsStore}
}
