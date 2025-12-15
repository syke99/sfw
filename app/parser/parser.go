package parser

import (
	"context"
	"fmt"
	extism "github.com/extism/go-sdk"

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

	states := make(map[string]map[string][]byte)

	for _, line := range web.Lines {
		state := make(map[string][]byte)

		states[line.Name] = state

		knots, ok := web.Knots[line.Name]
		if !ok {
			return nil, fmt.Errorf("knots for line %s not found", line.Name)
		}

		err := p.parseSecrets(knots, sw, state)
		if err != nil {
			// TODO: wrap error
			return nil, err
		}

		p.parseInputs(knots, sw)

		p.parseOutputs(knots, sw)
	}

	return sw, nil
}

func (p *parser) parseSecrets(knotGroup []*iModels.Knot, sw *models.StickyWeb, state map[string][]byte) error {
	scrts := make(map[string]string)
	manifests := make(map[string]*models.MessageManifest)

	for _, knot := range knotGroup {
		for _, key := range knot.Secrets {
			s, err := p.secretsStore.Get(key)
			if err != nil {
				return err
			}

			key = fmt.Sprintf("%s:%s", knot.Name, key)

			scrts[key] = s
		}

		// build extism config here
		manifests[knot.Name] = &models.MessageManifest{
			WasmPath:     knot.Knot,
			MemoryLimit:  knot.MemoryLimit,
			AllowedUrls:  knot.AllowedUrls,
			AllowedPaths: knot.AllowedPaths,
			HostFuncs:    composeHostFunctions(knot.Name, sw, state),
		}
	}

	sw.Secrets = scrts
	sw.Manifests = manifests

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

func composeHostFunctions(knotName string, sw *models.StickyWeb, state map[string][]byte) []extism.HostFunction {
	hostFuncs := make([]extism.HostFunction, 0)

	getSecrets := extism.NewHostFunctionWithStack(
		"get_secret",
		func(ctx context.Context, p *extism.CurrentPlugin, stack []uint64) {
			secretName, err := p.ReadString(stack[0])
			if err != nil {
				err = fmt.Errorf("error reading requested secret name: %w", err)

				stack[0], _ = p.WriteString(err.Error())
			}

			// TODO: trim name of plugin from secretName

			secret, ok := sw.Secrets[secretName]
			if !ok {
				stack[0], _ = p.WriteString("secret not found")
			}

			stack[0], err = p.WriteString(secret)
			if err != nil {
				stack[0], _ = p.WriteString("error writing back secret")
			}
		},
		[]extism.ValueType{extism.ValueTypePTR},
		[]extism.ValueType{extism.ValueTypePTR},
	)

	hostFuncs = append(hostFuncs, getSecrets)

	inputs := extism.NewHostFunctionWithStack(
		"get_input",
		func(ctx context.Context, p *extism.CurrentPlugin, stack []uint64) {
			inputName, err := p.ReadString(stack[0])
			if err != nil {
				err = fmt.Errorf("error reading requested input name: %w", err)

				stack[0], _ = p.WriteString(err.Error())
			}

			// TODO: remove knotName from inputName

			inputs, ok := sw.Inputs[inputName]
			if !ok {
				stack[0], _ = p.WriteString("input not found")
			}

			var iName string
			for _, input := range inputs {
				if input == inputName {
					iName = input
					break
				}
			}

			if iName == "" {
				stack[0], _ = p.WriteString("input not found")
				return
			}

			input, ok := state[iName]
			if !ok {
				stack[0], _ = p.WriteString("input not found")
			}

			stack[0], err = p.WriteBytes(input)
			if err != nil {
				stack[0], _ = p.WriteString("error writing back input")
			}
		},
		[]extism.ValueType{extism.ValueTypePTR},
		[]extism.ValueType{extism.ValueTypePTR},
	)

	hostFuncs = append(hostFuncs, inputs)

	outputs := extism.NewHostFunctionWithStack(
		"set_output",
		func(ctx context.Context, p *extism.CurrentPlugin, stack []uint64) {
			outputName, err := p.ReadString(stack[0])
			if err != nil {
				err = fmt.Errorf("error reading requested input name: %w", err)

				stack[0], _ = p.WriteString(err.Error())
			}

			outputValue, err := p.ReadBytes(stack[1])

			outputs, ok := sw.Outputs[outputName]
			if !ok {
				stack[0], _ = p.WriteString("input not found")
			}

			for _, output := range outputs {
				if output == outputName {
					if err != nil {
						stack[0], _ = p.WriteString("error getting output key")
					}
					state[outputName] = outputValue
					break
				}
			}
		},
		[]extism.ValueType{extism.ValueTypePTR, extism.ValueTypePTR},
		[]extism.ValueType{extism.ValueTypePTR},
	)

	hostFuncs = append(hostFuncs, outputs)

	return hostFuncs
}
