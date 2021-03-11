package cmd

import (
	"fmt"
	"os"

	"github.com/iiroj/go-git-staged/internal"
	"github.com/spf13/cobra"
)

// Execute go-lint-staged root command
func Execute(args []string) (failedCommands int) {
	// Declare variable for --all flag
	var allFiles bool
	// Get default working dir
	defaultWorkingDir, _ := os.Getwd()
	// Declare variable for --working-dir flag
	var workingDir string
	// Declare variable for --relative flag
	var relative bool
	// Declare variable for --verbose flag
	var verbose bool
	// Declare variable for --glob flags
	var globs []string
	// Declare variable for --commands flags
	var commands []string
	// Declare variable for number of failed commands
	failedCommands = 0

	// The main go-git-staged command
	var goGitStaged = &cobra.Command{
		Use:     "go-git-staged",
		Short:   "Run commands on files staged in git.\nFilter files using globs and pass them to their respective commands as arguments.",
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
		files, filesErr := internal.GetFiles(allFiles == false)
		if filesErr != nil {
			fmt.Printf("%s Failed to get staged files\n", internal.FailChar)
			return
		}
		filesLen := len(files)
		if filesLen == 0 {
			// Exit if there were no staged files
			fmt.Printf("%s No need to Go, working tree index is clean\n", internal.InfoChar)
			return
		} else if filesLen == 1 {
			// todo: is this the optimal way?
			fmt.Printf("%s Going with 1 staged file\n", internal.DoneChar)
		} else {
			// Update spinner with number of staged files
			fmt.Printf("%s Going with %d staged files\n", internal.DoneChar, filesLen)
		}

		// Parse --glob and --command args to a map with files
		globCommands, globCommandsErr := internal.ParseGlobCommands(args)
		if globCommandsErr != nil {
			fmt.Printf("%s %s", internal.FailChar, globCommandsErr.Error())
			return
		}

		// Normalize file paths to either absolute or relative to workingDir
		normalizedFiles, normalizedFilesErr := internal.NormalizeFiles(files, rootDir, relative, workingDir)
		if normalizedFilesErr != nil {
			fmt.Printf("%s %s", internal.FailChar, normalizedFilesErr.Error())
			return
		}

		// Match files to commands
		commands, _ := internal.MatchFilesToCommands(globCommands, normalizedFiles)

		// Create and run commands
		commandResults := internal.RunCommands(commands)

		// Gather up the number of failed results
		for _, commandResult := range commandResults {
			if commandResult.Err != nil {
				failedCommands++
			}
		}

		// Successful exit
		if failedCommands == 0 {
			fmt.Printf("%s Got git staged!\n", internal.DoneChar)

			// Print each command label and stdout
			if verbose == true {
				fmt.Println()

				for _, commandResult := range commandResults {
					if len(commandResult.Stdout) > 0 {
						fmt.Println(fmt.Sprintf("%s %s (%s):", internal.RunChar, commandResult.Label, commandResult.Info))
						fmt.Println(string(commandResult.Stdout))
					}
				}
			}

			return
		}

		// There were failed results
		if failedCommands == 1 {
			fmt.Printf("%s Got 1 failure from commands\n", internal.FailChar)
		} else {
			fmt.Printf("%s Got %d failes from commands:\n", internal.FailChar, failedCommands)
		}

		// Do not show cobra help message in this case
		// because we only report errors and call os.Exit(1)
		goGitStaged.SilenceUsage = true

		// Separate errors with an empty line
		fmt.Println()

		// Print the command label and error message for each fail
		for _, commandResult := range commandResults {
			if commandResult.Err != nil {
				fmt.Println(fmt.Sprintf("%s %s (%s):", internal.RunChar, commandResult.Label, commandResult.Info))
				fmt.Println(commandResult.Err.Error())
				fmt.Println()
			}
		}
	}

	// Do not sort flags, because --glob should come before --command
	goGitStaged.Flags().SortFlags = false

	// Add --all flag
	goGitStaged.Flags().BoolVarP(&allFiles, "all", "a", false, "Glob all files known to git instead of just staged")
	// Add --working-dir flag
	goGitStaged.Flags().StringVarP(&workingDir, "working-dir", "w", defaultWorkingDir, "Working directory for commands")
	// Add --relative flag
	goGitStaged.Flags().BoolVar(&relative, "relative", false, "Use file paths relative to --working-dir instead of absolute")
	// Add --verbose flag
	goGitStaged.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print command stdout after success instead of only fail")
	// Add --glob flags
	goGitStaged.Flags().StringArrayVarP(&globs, "glob", "g", globs, "Glob of files passed to following --command")
	goGitStaged.MarkFlagRequired("glob")
	// Add --commands flags
	goGitStaged.Flags().StringArrayVarP(&commands, "command", "c", commands, "Command to run with files matching previous --glob")
	goGitStaged.MarkFlagRequired("command")

	// If the execute failed for some other reason, assume 1 error
	if error := goGitStaged.Execute(); error != nil {
		return 1
	}

	// Return the number of failed commands
	return failedCommands
}
