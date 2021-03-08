package internal

import (
	"path/filepath"

	"github.com/pkg/errors"
)

// NormalizeFiles adjusts file paths to either absolute or relative to the given directory
func NormalizeFiles(files []string, basePath string, relative bool, relativeTo string) (normalizedFiles []string, err error) {
	normalizedFiles = make([]string, len(files))

	// Resolve absolute paths for files relative to basePath
	if relative == false {
		for i, file := range files {
			normalizedFiles[i] = filepath.Join(basePath, file)
		}

		return normalizedFiles, nil
	}

	for i, file := range files {
		relativeFile, relativeFileError := filepath.Rel(relativeTo, filepath.Join(basePath, file))
		if relativeFileError != nil {
			return nil, errors.New("Failed to normalize relative filenames")
		}
		normalizedFiles[i] = relativeFile
	}

	return normalizedFiles, nil
}
