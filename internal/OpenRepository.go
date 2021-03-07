package internal

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

// findDotGit returns the first directory containing .git, starting from the given path
func findDotGit(start string) (directory string) {
	root := string(os.PathSeparator)

	for {
		if abs, _ := filepath.Abs(start); abs == root {
			// Stop at chroot
			break
		}

		if _, statError := os.Stat(filepath.Join(start, ".git")); statError != nil {
			// If .git doesn't exist, go up one level
			start += "/.."
		} else {
			// Otherwise, break out of loop
			break
		}
	}

	return start
}

// OpenRepository resolves the git repository from the supplied pathname
func OpenRepository(pathname string) (repository *git.Repository, repositoryRoot string, err error) {
	// Fail if pathname does not exist
	if _, statError := os.Stat(pathname); os.IsNotExist(statError) {
		return nil, "", statError
	}

	// Find directory containing .git, starting from pathname
	directory := findDotGit(pathname)

	// Open repository
	repository, repositoryError := git.PlainOpen(directory)
	if repositoryError != nil {
		return nil, "", repositoryError
	}

	return repository, directory, nil
}
