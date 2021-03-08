package internal

import (
	"fmt"
	"sync"
	"time"

	"github.com/theckman/yacspin"
)

// Some configuration for spinner.
// Characters contain a space to separate from messages.
var spinnerConfig = yacspin.Config{
	Frequency:         100 * time.Millisecond,
	CharSet:           yacspin.CharSets[14],
	StopCharacter:     "  ✓ ",
	StopColors:        []string{"fgGreen"},
	StopFailCharacter: "  𐄂 ",
	StopFailColors:    []string{"fgRed"},
}

// RunCommands creates commands and runs them
func RunCommands(commands []Command) (err error) {
	var wg sync.WaitGroup

	for _, command := range commands {
		wg.Add(1)
		globs := command.globs
		filesLen := len(command.files)
		commands := command.commands

		go func() {
			defer wg.Done()
			spinner, spinnerErr := yacspin.New(spinnerConfig)
			if spinnerErr == nil {
				spinner.Start()
				spinner.StopMessage(fmt.Sprintf("%s (%d files): %s", globs, filesLen, commands))
				spinner.Stop()
			}
		}()
	}

	wg.Wait()

	return nil
}
