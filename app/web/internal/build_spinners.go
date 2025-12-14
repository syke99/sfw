package internal

import (
	"fmt"
	"github.com/syke99/fn"
	"github.com/syke99/sfw/app/parser"
	"github.com/syke99/sfw/app/secrets"
	"github.com/syke99/sfw/app/spinner"
	"github.com/syke99/sfw/pkg/models"
)

func BuildSpinners(path string) (map[string]spinner.Spinner, error) {
	w := fn.Try(buildWeb, path, nil)

	spinners, err := fn.Try(func(w *models.Web) (map[string]spinner.Spinner, error) {
		scrts := secrets.NewSecretsStore(w)

		psr := parser.NewParser(scrts)

		spns := make(map[string]spinner.Spinner)

		for _, line := range w.Lines {
			if line.Trigger.Webhook != "" {
				sp, err := spinner.NewWebSpinner(w, psr)
				if err != nil {
					err = fmt.Errorf("error creating webhook spinner: %w", err)
					return nil, err
				}
				spns[line.Trigger.Webhook] = sp
				continue
			}

			sp, err := spinner.NewFileSpinner(w, psr)
			if err != nil {
				err = fmt.Errorf("error creating file spinner: %w", err)
				return nil, err
			}
			spns[line.Trigger.File] = sp
		}

		return spns, nil
	}, w, nil).Out()
	if err != nil {
		return nil, err
	}

	return *spinners, nil
}
