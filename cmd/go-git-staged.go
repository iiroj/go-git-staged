package cmd

import (
	"fmt"
	"os"

	"github.com/iiroj/go-git-staged/internal"
	"github.com/spf13/cobra"
)

// Execute go-lint-staged root command
func Execute(args []string) (failedCommands int) {
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
	// Declare variable for number of failed commands
	failedCommands = 0

	// The main go-git-staged command
	var goGitStaged = &cobra.Command{
		Use:     "go-git-staged",
		Short:   "Run commands for files staged in git",
		Example: "go-git-staged --glob '*.js' --command 'eslint' --command 'prettier'",
	}

	goGitStaged.Run = func(cmd *cobra.Command, _ []string) {
		// Resolve git root directory
		rootDir, rootDirErr := internal.ResolveRootDir(workingDir)
		if rootDirErr != nil {
			fmt.Printf("%s Failed to resolve git directory\n", internal.FailChar)
			return
		}

		// Get staged files
		stagedFiles, stagedFilesError := internal.GetStagedFiles()
		if stagedFilesError != nil {
			fmt.Printf("%s Failed to get staged files\n", internal.FailChar)
			return
		}
		stagedFilesLen := len(stagedFiles)
		if stagedFilesLen == 0 {
			// Exit if there were no staged files
			fmt.Printf("%s No need to Go, working tree index is clean\n", internal.InfoChar)
			return
		} else if stagedFilesLen == 1 {
			// todo: is this the optimal way?
			fmt.Printf("%s Going with 1 staged file\n", internal.DoneChar)
		} else {
			// Update spinner with number of staged files
			fmt.Printf("%s Going with %d staged files\n", internal.DoneChar, len(stagedFiles))
		}

		// Parse --glob and --command args to a map with files
		globCommands, globCommandsErr := internal.ParseGlobCommands(args)
		if globCommandsErr != nil {
			fmt.Printf("%s %s", internal.FailChar, globCommandsErr.Error())
			return
		}

		// Normalize file paths to either absolute or relative to workingDir
		normalizedFiles, normalizedFilesErr := internal.NormalizeFiles(stagedFiles, rootDir, relative, workingDir)
		if normalizedFilesErr != nil {
			fmt.Printf("%s %s", internal.FailChar, normalizedFilesErr.Error())
			return
		}

		// Match files to commands
		commands, _ := internal.MatchFilesToCommands(globCommands, normalizedFiles)

		// Create and run commands
		commandResults := internal.RunCommands(commands)

		for _, commandResult := range commandResults {
			if commandResult.Err != nil {
				failedCommands++
			}
		}

		if failedCommands == 0 {
			fmt.Printf("%s Got git staged!\n", internal.DoneChar)
			return
		}

		if failedCommands == 1 {
			fmt.Printf("%s Got 1 failure from commands\n", internal.FailChar)
		} else {
			fmt.Printf("%s Got %d failes from commands:\n", internal.FailChar, failedCommands)
		}

		goGitStaged.SilenceUsage = true

		fmt.Println()

		for _, commandResult := range commandResults {
			if commandResult.Err != nil {
				fmt.Println(fmt.Sprintf("%s %s:", internal.FailChar, commandResult.Label))
				fmt.Println(commandResult.Err.Error())
				fmt.Println()
			}
		}
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
		return 1
	}

	return failedCommands
}
