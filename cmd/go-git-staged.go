package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/iiroj/go-git-staged/internal"
	"github.com/spf13/cobra"
	"github.com/theckman/yacspin"
)

// Some configuration for spinner.
// Characters contain a space to separate from messages.
var spinnerConfig = yacspin.Config{
	Frequency:         100 * time.Millisecond,
	CharSet:           yacspin.CharSets[14],
	StopCharacter:     "✓ ",
	StopColors:        []string{"fgGreen"},
	StopFailCharacter: "𐄂 ",
	StopFailColors:    []string{"fgRed"},
}

// Result for Execute
type Result struct {
	Error error
}

// Execute main entrypoint
func Execute(args []string) Result {
	// Declare result for error handling
	var result Result

	// Create spinner with static config from above
	spinner, spinnerError := yacspin.New(spinnerConfig)
	// Spinner error handling
	if spinnerError != nil {
		result.Error = spinnerError
		return result
	}

	// Get default working dir
	defaultWorkingDir, _ := os.Getwd()
	// Declare variable for --working-dir flag
	var workingDir string
	// Declare variable for --relative flag
	var relative bool

	// The main go-git-staged command
	var goGitStaged = &cobra.Command{
		Use:   "go-git-staged",
		Short: "Run commands for files staged in git",
		Run: func(cmd *cobra.Command, args []string) {
			// Start spinner
			spinner.Start()

			// Open git repository
			repository, repositoryRoot, repositoryError := internal.OpenRepository(workingDir)
			if repositoryError != nil {
				spinner.StopFailMessage("Failed to open git repository")
				spinner.StopFail()
				return
			}

			// Get staged files
			stagedFiles, stagedFilesError := internal.GetStagedFiles(repository)
			if stagedFilesError != nil {
				spinner.StopFailMessage("Failed to get staged files")
				spinner.StopFail()
				return
			}

			stagedFilesLen := len(stagedFiles)

			if stagedFilesLen == 0 {
				// Exit if there were no staged files
				spinner.StopCharacter("ℹ︎ ")
				spinner.StopColors("fgBlue")
				spinner.StopMessage("No need to Go, working tree index is clean")
			} else if stagedFilesLen == 1 {
				// todo: is this the optimal way?
				spinner.StopMessage("Going with 1 staged file")
			} else {
				// Update spinner with number of staged files
				spinner.StopMessage(fmt.Sprintf("Going with %d staged files", len(stagedFiles)))
			}

			filepaths := internal.NormalizeFiles(stagedFiles, repositoryRoot, relative, workingDir)

			fmt.Println(filepaths)

			spinner.Stop()
			return
		},
	}

	// Add a persistent --working-dir flag
	goGitStaged.PersistentFlags().StringVarP(&workingDir, "working-dir", "w", defaultWorkingDir, "Working directory for commands")
	// Add --relative flag
	goGitStaged.Flags().BoolVar(&relative, "relative", false, "Use file paths relative to --working-dir (default \"false\")")

	// Handle error by returning it in result
	Error := goGitStaged.Execute()
	if Error != nil {
		result.Error = Error
		return result
	}

	return result
}
