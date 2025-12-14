package internal

import (
	"fmt"
	"github.com/syke99/fn"
	"github.com/syke99/sfw/app/parser"
	"github.com/syke99/sfw/app/secrets"
	"github.com/syke99/sfw/app/spinner"
	"github.com/syke99/sfw/pkg/models"
)

func BuildSpinners(path string) ([]spinner.Spinner, error) {
	w := fn.Try(buildWeb, path, nil)

	spinners, err := fn.Try(func(w *models.Web) ([]spinner.Spinner, error) {
		scrts := secrets.NewSecretsStore(w)

		psr := parser.NewParser(scrts)

		spns := make([]spinner.Spinner, len(w.Lines))

		for i, line := range w.Lines {
			if line.Trigger.Webhook != "" {
				sp, err := spinner.NewWebSpinner(w, psr, line.Trigger.Webhook)
				if err != nil {
					err = fmt.Errorf("error creating webhook spinner: %w", err)
					return nil, err
				}
				spns[i] = sp
				continue
			}

			sp, err := spinner.NewFileSpinner(w, psr, line.Trigger.File)
			if err != nil {
				err = fmt.Errorf("error creating file spinner: %w", err)
				return nil, err
			}
			spns[i] = sp
		}

		return spns, nil
	}, w, nil).Out()
	if err != nil {
		return nil, err
	}

	return *spinners, nil
}
