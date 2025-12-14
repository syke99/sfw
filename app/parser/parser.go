package parser

import (
	"github.com/syke99/sfw/app/secrets"
	"github.com/syke99/sfw/pkg/models"
)

type parser struct {
	secretsStore secrets.SecretsStore
}

func (p parser) Parse(web *models.Web) (*StickyWeb, error) {
	sw := &StickyWeb{}

	err := p.parseSecrets(sw)

	// TODO: implement inputs and outputs
	// TODO: and add

	if err != nil {
		// TODO: wrap error
		return nil, err
	}

	return sw, nil
}

func (p parser) parseSecrets(sw *StickyWeb) error {
	//TODO implement parsing here
	return nil
}

func NewParser(secretsStore secrets.SecretsStore) Parser {
	return &parser{secretsStore: secretsStore}
}
