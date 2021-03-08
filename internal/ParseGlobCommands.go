package internal

import (
	"github.com/pkg/errors"
)

// Pair holds mapping of globs and their commands
type Pair struct {
	globs, commands []string
}

// ParseGlobCommands parses the command line arguments --glob and --command to a mapping
func ParseGlobCommands(args []string) (pairs []Pair, err error) {
	// Filter args other than --working-dir or --relative
	filteredArgs := make([]string, 0)
	for _, str := range args {
		if str != "-w" && str != "--working-dir" && str != "--relative" {
			filteredArgs = append(filteredArgs, str)
		}
	}

	// Initialize globCommands
	pairs = make([]Pair, 0)
	// At the start, current arg hast to be a glob
	currentIsGlob, currentGlobIndex := true, 0

	for i, str := range filteredArgs {
		// If we're dealing with the --glob flag itself
		if str == "-g" || str == "--glob" {
			// If the arg is not a glob, it was a command
			// This new --glob flag marks a second group of GlobCommands
			if currentIsGlob == false {
				currentGlobIndex++
				currentIsGlob = true
			}

			continue
		} else if i == 0 {
			// Fail if the first flag was a --command
			return nil, errors.New("--command should come after at least one --glob")
		}

		// If the current flag is a --command, the following args are command strings
		if str == "-c" || str == "--command" {
			currentIsGlob = false
			continue
		}

		if currentIsGlob {
			if len(pairs) > currentGlobIndex {
				globs := pairs[currentGlobIndex].globs
				commands := pairs[currentGlobIndex].commands
				// Append arg to the current globs
				pairs[currentGlobIndex] = Pair{append(globs, str), commands}
			} else {
				globs := make([]string, 1)
				globs = append(globs, str)
				commands := make([]string, 0)
				pairs = append(pairs, Pair{globs, commands})
			}
		} else if len(pairs) > currentGlobIndex {
			globs := pairs[currentGlobIndex].globs
			commands := pairs[currentGlobIndex].commands
			// Append arg to the current commands
			pairs[currentGlobIndex] = Pair{globs, append(commands, str)}
		}
	}

	for _, pair := range pairs {
		if len(pair.commands) == 0 {
			return nil, errors.Errorf("Got no commands for glob %s", pair.globs)
		}
	}

	return pairs, nil
}
