package internal

import (
	"os"
	"os/exec"
)

// ResolveRootDir resolves the git root repository containing .git from the supplied pathname
func ResolveRootDir(pathname string) (rootDir string, err error) {
	// Fail if pathname does not exist
	if _, statError := os.Stat(pathname); os.IsNotExist(statError) {
		return "", statError
	}

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
