package internal

import (
	"path/filepath"

	"github.com/pkg/errors"
)

// NormalizeFiles adjusts file paths to either absolute or relative to the given directory
func NormalizeFiles(files []string, basePath string, relative bool, relativeTo string) (normalizedFiles []string, err error) {
	// Resolve absolute paths for files relative to basePath
	absoluteFiles := make([]string, len(files))
	for i, file := range files {
		absoluteFiles[i] = filepath.Join(basePath, file)
	}

	if !relative {
		return absoluteFiles, nil
	}

	for i, file := range absoluteFiles {
		// Relative to the current CWD, instead of the git repo root
		relativeFile, relativeFileError := filepath.Rel(relativeTo, file)
		if relativeFileError != nil {
			return nil, errors.New("Failed to normalize relative filenames")
		}
		normalizedFiles[i] = relativeFile
	}

	return normalizedFiles, nil
}
