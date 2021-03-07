package internal

import (
	"path/filepath"
)

// NormalizeFiles adjusts file paths to either absolute or relative to the given directory
func NormalizeFiles(files []string, basePath string, relative bool, relativeTo string) (normalizedFiles []string) {
	// Resolve absolute paths for files relative to basePath
	if relative == false {
		absoluteFiles := make([]string, len(files))
		for _, file := range files {
			absoluteFile := filepath.Join(basePath, file)
			absoluteFiles = append(absoluteFiles, absoluteFile)
		}

		return absoluteFiles
	}

	relativePaths := make([]string, len(files))
	for _, file := range files {
		relativeFile, relativeFileError := filepath.Rel(relativeTo, filepath.Join(basePath, file))
		if relativeFileError != nil {
			break
		}
		relativePaths = append(relativePaths, relativeFile)
	}

	return relativePaths
}
