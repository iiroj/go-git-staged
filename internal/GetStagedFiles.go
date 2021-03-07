package internal

import (
	"github.com/go-git/go-git/v5"
)

// GetStagedFiles returns a list of staged files of the given git repository
func GetStagedFiles(repository *git.Repository) (files []string, err error) {
	// Open worktree
	worktree, worktreeError := repository.Worktree()
	if worktreeError != nil {
		return nil, worktreeError
	}

	// Get worktree status
	status, statusError := worktree.Status()
	if statusError != nil {
		return nil, statusError
	}

	// Initialize staged files array
	stagedFiles := make([]string, 0)

	// Iterate statuses
	for filename, fileStatus := range status {
		// Add if file is added, copied, modified, or renamed in the staging area
		switch fileStatus.Staging {
		case git.Added, git.Copied, git.Modified, git.Renamed:
			stagedFiles = append(stagedFiles, filename)
		}
	}

	return stagedFiles, nil
}
