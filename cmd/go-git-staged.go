package cmd

import (
	"fmt"
	"os"
	"time"

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

	// The main go-git-staged command
	var goGitStaged = &cobra.Command{
		Use:   "go-git-staged",
		Short: "Run commands for files staged in git",
		Run: func(cmd *cobra.Command, args []string) {
			// Start spinner
			spinner.Start()

			// Check if workingDir exists
			_, statError := os.Stat(workingDir)
			// Fail if workingDir does not exist
			if os.IsNotExist(statError) {
				spinner.StopFailMessage("Failed to find a working directory")
				spinner.StopFail()
				return
			}

			// Print workingDir and stop spinner.
			// This is all the functionality for now.
			spinner.StopMessage(fmt.Sprintf("workingDir: %s", workingDir))
			spinner.Stop()
		},
	}

	// Add a persistent --working-dir flag
	goGitStaged.PersistentFlags().StringVarP(&workingDir, "working-dir", "w", defaultWorkingDir, "Working directory for commands")

	// Handle error by returning it in result
	Error := goGitStaged.Execute()
	if Error != nil {
		result.Error = Error
		return result
	}

	return result
}
