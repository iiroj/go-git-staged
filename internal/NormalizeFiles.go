package internal

import (
	"path/filepath"

	"github.com/pkg/errors"
)

// NormalizeFiles adjusts file paths to either absolute or relative to the given directory
func NormalizeFiles(files []string, basePath string, relative bool, relativeTo string) (normalizedFiles []string, err error) {
	// Resolve absolute paths for files relative to basePath
	if relative == false {
		absoluteFiles := make([]string, len(files))
		for i, file := range files {
			absoluteFiles[i] = filepath.Join(basePath, file)
		}

		return absoluteFiles, nil
	}

	relativeFiles := make([]string, len(files))
	for i, file := range files {
		relativeFile, relativeFileError := filepath.Rel(relativeTo, filepath.Join(basePath, file))
		if relativeFileError != nil {
			return nil, errors.New("Failed to normalize relative filenames")
		}
		relativeFiles[i] = relativeFile
	}

	return relativeFiles, nil
}
