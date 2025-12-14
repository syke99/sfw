package internal

import (
	"github.com/syke99/fn"
	iModels "github.com/syke99/sfw/internal/pkg/models"
	"github.com/syke99/sfw/pkg"
	"github.com/syke99/sfw/pkg/models"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

type ptr struct {
	web  *models.Web
	path string
}

func buildWeb(path string) (*models.Web, error) {
	_web := &models.Web{}

	webFilePath := filepath.Join(path, "web.yaml")

	// TODO: wrapper error
	webFile := fn.Try(os.Open, webFilePath, pkg.OpeningWebFileError)

	data := fn.Try(io.ReadAll, webFile, pkg.ReadWebFileError)

	w := new(iModels.Web)

	_, err := fn.Try(func(yamlData []byte) (interface{}, error) {
		err := yaml.Unmarshal(yamlData, w)

		return nil, err
	}, data, pkg.WebUnmarshalError).Out()
	if err != nil {
		return nil, err
	}

	_web.Web = w

	_, err = fn.Try(buildLines, &ptr{
		web:  _web,
		path: webFilePath,
	}, pkg.BuildLinesError).Out()
	if err != nil {
		return nil, err
	}

	return _web, nil
}
