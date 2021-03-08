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

// Execute go-lint-staged root command
func Execute(args []string) (err error) {
	// Create spinner with static config from above
	spinner, spinnerError := yacspin.New(spinnerConfig)
	// Spinner error handling
	if spinnerError != nil {
		return spinnerError
	}

	// Get default working dir
	defaultWorkingDir, _ := os.Getwd()
	// Declare variable for --working-dir flag
	var workingDir string
	// Declare variable for --relative flag
	var relative bool
	// Declare variable for --glob flags
	var globs []string
	// Declare variable for --commands flags
	var commands []string

	// The main go-git-staged command
	var goGitStaged = &cobra.Command{
		Use:     "go-git-staged",
		Short:   "Run commands for files staged in git",
		Example: "go-git-staged --glob '*.js' --command 'eslint' --command 'prettier'",
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			// Start spinner
			spinner.Start()

			// Open git repository
			repository, repositoryRoot, repositoryErr := internal.OpenRepository(workingDir)
			if repositoryErr != nil {
				spinner.StopFailMessage("Failed to open git repository")
				spinner.StopFail()
				return repositoryErr
			}

			// Get staged files
			stagedFiles, stagedFilesError := internal.GetStagedFiles(repository)
			if stagedFilesError != nil {
				spinner.StopFailMessage("Failed to get staged files")
				spinner.StopFail()
				return stagedFilesError
			}
			stagedFilesLen := len(stagedFiles)
			if stagedFilesLen == 0 {
				// Exit if there were no staged files
				spinner.StopCharacter("ℹ︎ ")
				spinner.StopColors("fgBlue")
				spinner.StopMessage("No need to Go, working tree index is clean")
				spinner.Stop()
				return nil
			} else if stagedFilesLen == 1 {
				// todo: is this the optimal way?
				spinner.StopMessage("Going with 1 staged file")
			} else {
				// Update spinner with number of staged files
				spinner.StopMessage(fmt.Sprintf("Going with %d staged files", len(stagedFiles)))
			}

			spinner.Stop()
			spinner.Start()

			// Parse --glob and --command args to a map with files
			globCommands, globCommandsErr := internal.ParseGlobCommands(args)
			if globCommandsErr != nil {
				spinner.StopFailMessage(globCommandsErr.Error())
				spinner.StopFail()
				return globCommandsErr
			}
			spinner.Message("Got a valid configuration")

			// Normalize file paths to either absolute or relative to workingDir
			normalizedFiles, normalizedFilesErr := internal.NormalizeFiles(stagedFiles, repositoryRoot, relative, workingDir)
			if normalizedFilesErr != nil {
				spinner.StopFailMessage(normalizedFilesErr.Error())
				spinner.StopFail()
				return normalizedFilesErr
			}
			if relative == true {
				spinner.Message("Got relative filenames")
			} else {
				spinner.Message("Got absolute filenames")
			}

			// Match files to commands
			commands, commandsErr := internal.MatchFilesToCommands(globCommands, normalizedFiles)
			if commandsErr != nil {
				spinner.StopFailMessage(commandsErr.Error())
				spinner.StopFail()
				return commandsErr
			}

			// Create and run commands
			if runErr := internal.RunCommands(commands); err != nil {
				spinner.StopFailMessage(runErr.Error())
				spinner.StopFail()
				return runErr
			}

			spinner.StopMessage("Got git staged!")
			spinner.Stop()
			return nil
		},
	}

	// Do not sort flags, because --glob should come before --command
	goGitStaged.Flags().SortFlags = false

	// Add --working-dir flag
	goGitStaged.Flags().StringVarP(&workingDir, "working-dir", "w", defaultWorkingDir, "Working directory for commands")
	// Add --relative flag
	goGitStaged.Flags().BoolVar(&relative, "relative", false, "Use file paths relative to --working-dir (default \"false\")")
	// Add --glob flags
	goGitStaged.Flags().StringArrayVarP(&globs, "glob", "g", globs, "Glob of files passed to following --command")
	goGitStaged.MarkFlagRequired("glob")
	// Add --commands flags
	goGitStaged.Flags().StringArrayVarP(&commands, "command", "c", commands, "Command to run with files matching previous --glob")
	goGitStaged.MarkFlagRequired("command")

	// Handle error by returning it in result
	if error := goGitStaged.Execute(); error != nil {
		return error
	}

	return nil
}
