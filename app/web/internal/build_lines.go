package internal

import (
	"github.com/syke99/fn"
	iModels "github.com/syke99/sfw/internal/pkg/models"
	"github.com/syke99/sfw/pkg"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

func buildLines(ptr *ptr) (interface{}, error) {
	linesPath := filepath.Join(ptr.path, "lines")

	linesDir := fn.Try(os.Open, linesPath, pkg.OpenLinesDirError)

	linesEntries := fn.Try(func(dir *os.File) ([]os.DirEntry, error) {
		return dir.ReadDir(0)
	}, linesDir, pkg.ReadLinesEntriesError)

	lines, err := fn.Try(func(in []os.DirEntry) ([]*iModels.Line, error) {
		l := make([]*iModels.Line, len(in))

		for i, line := range in {
			entry := new(iModels.Line)

			entryPath := filepath.Join(linesPath, line.Name())

			lineFile := fn.Try(os.Open, entryPath, pkg.ReadLineError)

			lineData := fn.Try(io.ReadAll, lineFile, nil)

			_, err := fn.Try(func(yamlData []byte) (interface{}, error) {
				err := yaml.Unmarshal(yamlData, entry)

				return nil, err
			}, lineData, pkg.LineUnmarshalError).Out()
			if err != nil {
				return nil, err
			}

			l[i] = entry

			ptr.path = entryPath

			knots, err := fn.Try(buildKnots, ptr, pkg.BuildKnotsError).Out()
			if err != nil {
				return nil, err
			}

			ptr.web.Knots[entry.Name] = *knots
		}

		return l, nil
	}, linesEntries, pkg.BuildLinesError).Out()
	if err != nil {
		return nil, err
	}

	ptr.web.Lines = *lines

	return nil, nil
}
