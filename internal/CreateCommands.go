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
	var wg sync.WaitGroup
	commandResults = make([]CommandResult, 0)

	for _, command := range commands {
		files := command.files

		wg.Add(1)

		fmt.Printf("  %s %d files (%s):\n", RunChar, len(files), strings.Join(command.globs, ", "))

		for _, command := range command.commands {
			fmt.Printf("    %s\n", command)

			go func(command string) {
				cmd := exec.Command(command, files...)
				_, err := cmd.Output()
				commandResults = append(commandResults, CommandResult{command, err})
				defer wg.Done()
			}(command)
		}
	}

	wg.Wait()

	return commandResults
}
