package internal

import (
	"os"
	"os/exec"
	"path/filepath"
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

// ResolveRootDir resolves the git root repository containing .git from the supplied pathname
func ResolveRootDir(pathname string) (rootDir string, err error) {
	// Fail if pathname does not exist
	if _, statError := os.Stat(pathname); os.IsNotExist(statError) {
		return "", statError
	}

	// Find directory containing .git, starting from pathname
	rootDir = findDotGit(pathname)

	// Get output directly from git
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	rootDir = string(stdout)
	rootDir = rootDir[:len(rootDir)-1] // Remove linebreak

	return rootDir, nil
}
