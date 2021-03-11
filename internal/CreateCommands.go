package internal

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

// CommandResult contains the result of a single command
type CommandResult struct {
	Label string
	Err   error
}

// RunCommands creates commands and runs them
func RunCommands(commands []Command) (commandResults []CommandResult) {
	// Use a WaitGroup for all goroutines
	var wg sync.WaitGroup
	// Initialize slice for CommandResults
	commandResults = make([]CommandResult, 0)

	for _, command := range commands {
		files := command.files

		// Print a label for the current glob slice, and how many files they matched
		fmt.Printf("  %s %d files (%s):\n", RunChar, len(files), strings.Join(command.globs, ", "))

		for _, command := range command.commands {
			// Print the current command as a string
			fmt.Printf("    %s\n", command)

			wg.Add(1)

			// Run the command in a goroutine, and gather errors
			go func(command string) {
				cmd := exec.Command(command, files...)
				_, err := cmd.Output()
				commandResults = append(commandResults, CommandResult{command, err})
				wg.Done()
			}(command)
		}
	}

	// Wait until all commands are done
	wg.Wait()

	return commandResults
}
