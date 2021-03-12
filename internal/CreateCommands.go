package internal

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/iiroj/go-git-staged/char"
)

// CommandResult contains the result of a single command
type CommandResult struct {
	Label  string
	Info   string
	Stdout []byte
	Err    error
}

// RunCommands creates commands and runs them
func RunCommands(commands []Command) (commandResults []CommandResult) {
	// Use a WaitGroup for all goroutines
	var wg sync.WaitGroup
	// Initialize slice for CommandResults
	commandResults = make([]CommandResult, 0)
	// Initialize slice for all command groups.
	// The inner list should be run serially, and the outer concurrently.
	commandGroups := make([][]func(), len(commands))

	for i, command := range commands {
		files := command.files

		// Print a label for the current glob slice, and how many files they matched
		info := fmt.Sprintf("%d files (%s)", len(files), strings.Join(command.globs, ", "))
		fmt.Printf("  %s %s:\n", char.Run, info)

		// Group all commands
		commandGroup := make([]func(), len(command.commands))

		for k, c := range command.commands {
			// Assign loop variable so it stays in the func
			command := c
			// Print the current command as a string
			fmt.Printf("    %s\n", command)

			// Run the command in a goroutine, and gather errors
			commandGroup[k] = func() {
				cmd := exec.Command(command, files...)
				stdout, err := cmd.Output()
				commandResults = append(commandResults, CommandResult{command, info, stdout, err})
			}
		}

		commandGroups[i] = commandGroup
	}

	wg.Add(len(commandGroups))
	for _, commands := range commandGroups {
		runCommands := func(commands []func()) {
			// Run all commands in a group serially
			for _, command := range commands {
				command()
			}
			wg.Done()
		}

		// Run all groups concurrently
		go runCommands(commands)
	}

	// Wait until all commands are done
	wg.Wait()

	return commandResults
}
