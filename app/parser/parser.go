package parser

import (
	"fmt"
	"github.com/syke99/sfw/app/secrets"
	iModels "github.com/syke99/sfw/internal/pkg/models"
	"github.com/syke99/sfw/pkg/models"
)

type parser struct {
	secretsStore secrets.SecretsStore
}

func (p *parser) Parse(web *models.Web) (*models.StickyWeb, error) {
	sw := new(models.StickyWeb)

	sw.Web = web

	for _, knots := range web.Knots {
		err := p.parseSecrets(knots, sw)
		if err != nil {
			// TODO: wrap error
			return nil, err
		}

		p.parseInputs(knots, sw)

		p.parseOutputs(knots, sw)
	}

	return sw, nil
}

func (p *parser) parseSecrets(knotGroup []*iModels.Knot, sw *models.StickyWeb) error {
	scrts := make(map[string]string)

	for _, knot := range knotGroup {
		for _, key := range knot.Secrets {
			s, err := p.secretsStore.Get(key)
			if err != nil {
				return err
			}

			key = fmt.Sprintf("%s:%s", knot.Name, key)

			scrts[key] = s
		}
	}

	sw.Secrets = scrts

	return nil
}

func (p *parser) parseInputs(knotGroup []*iModels.Knot, sw *models.StickyWeb) {
	inputs := make(map[string][]string)

	for _, knot := range knotGroup {
		for _, key := range knot.Inputs {
			key = fmt.Sprintf("%s:%s", knot.Name, key)

			inputs[key] = append(inputs[key], key)
		}
	}

	sw.Inputs = inputs
}

func (p *parser) parseOutputs(knotGroup []*iModels.Knot, sw *models.StickyWeb) {
	outputs := make(map[string][]string)

	for _, knot := range knotGroup {
		for _, key := range knot.Outputs {
			key = fmt.Sprintf("%s:%s", knot.Name, key)

			outputs[key] = append(outputs[key], key)
		}
	}

	sw.Outputs = outputs
}

func NewParser(secretsStore secrets.SecretsStore) Parser {
	return &parser{secretsStore: secretsStore}
}
