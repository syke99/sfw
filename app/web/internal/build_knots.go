package internal

import (
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/syke99/fn"
	iModels "github.com/syke99/sfw/internal/pkg/models"
	"github.com/syke99/sfw/pkg"
)

func buildKnots(ptr *ptr) ([]*iModels.Knot, error) {
	knotsPath := filepath.Join(ptr.path, "knots")

	knotsDir := fn.Try(os.Open, knotsPath, pkg.OpenKnotsDirError)

	knotsEntries := fn.Try(func(dir *os.File) ([]os.DirEntry, error) {
		return dir.ReadDir(0)
	}, knotsDir, pkg.ReadKnotsEntriesError)

	knots, err := fn.Try(func(in []os.DirEntry) ([]*iModels.Knot, error) {
		knots := make([]*iModels.Knot, len(in))

		for i, knot := range in {
			entry := new(iModels.Knot)

			entryPath := filepath.Join(knotsPath, knot.Name())

			knotFile := fn.Try(os.Open, entryPath, pkg.OpenKnotsDirError)

			knotData := fn.Try(io.ReadAll, knotFile, pkg.ReadKnotError)

			_, err := fn.Try(func(yamlData []byte) (interface{}, error) {
				err := yaml.Unmarshal(yamlData, entry)

				return nil, err
			}, knotData, pkg.KnotUnmarshalError).Out()
			if err != nil {
				return nil, err
			}

			knots[i] = entry
		}

		return knots, nil
	}, knotsEntries, nil).Out()
	if err != nil {
		return nil, err
	}

	return *knots, nil
}
