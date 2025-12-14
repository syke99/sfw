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
		spns := make([]spinner.Spinner, len(w.Lines))

		vault := w.Web.Secrets.Vault

		scrts, err := secrets.NewSecretsStore(vault.Mount, vault.URL, vault.SecretID, vault.RoleID)
		if err != nil {
			return nil, err
		}

		psr := parser.NewParser(scrts)

		stickyWeb, err := psr.Parse(w)
		if err != nil {
			return nil, err
		}

		for i, line := range w.Lines {
			if line.Trigger.Webhook != "" {
				sp, err := spinner.NewWebSpinner(w, stickyWeb, line.Trigger.Webhook)
				if err != nil {
					err = fmt.Errorf("error creating webhook spinner: %w", err)
					return nil, err
				}
				spns[i] = sp
				continue
			}

			sp, err := spinner.NewFileSpinner(w, stickyWeb, line.Trigger.File)
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
