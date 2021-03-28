package internal

import (
	"os/exec"
	"strings"
)

// GetFiles returns a list of files in the git index, or optionally only staged
func GetFiles(stagedOnly bool) (files []string, err error) {
	var cmd *exec.Cmd

	if stagedOnly {
		// The -z flag makes sure files are unquoted and separated by \u0000
		// See https://git-scm.com/docs/git-diff#Documentation/git-diff.txt--z
		cmd = exec.Command("git", "diff", "--staged", "--diff-filter=ACMR", "--name-only", "-z")
	} else {
		cmd = exec.Command("git", "ls-files", "-z")
	}

	stdout, err := cmd.Output()
	if err != nil {
		return files, err
	}

	// Split string form \u0000 to get slice of files
	files = strings.Split(string(stdout), "\u0000")
	// The -z flags leaves a \u0000 at the end, so remove the last empty item
	files = files[:len(files)-1]

	return files, nil
}
