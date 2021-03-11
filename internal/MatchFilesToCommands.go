package internal

import (
	"github.com/bmatcuk/doublestar/v2"
)

// Command keeps a mapping of filenames and their commands
type Command struct {
	globs, files, commands []string
}

// MatchFilesToCommands resolves files to their commands
func MatchFilesToCommands(pairs []Pair, files []string) (commands []Command, err error) {
	commands = make([]Command, 0)

	// Loop through all glob, command pairs
	for _, pair := range pairs {
		// Gather up matched files
		matches := make([]string, 0)
		// Loop through all globs in the current pair
		for _, glob := range pair.globs {
			for _, filename := range files {
				match, matchError := doublestar.PathMatch(glob, filename)

				// If the glob is not valid, exit with an error
				if matchError != nil {
					return nil, matchError
				}

				// If the filename matched the glob, add matches
				if match == true {
					matches = append(matches, filename)
				}
			}
		}

		// Add a result if there were any matches, no need to otherwise
		if len(matches) > 0 {
			commands = append(commands, Command{pair.globs, matches, pair.commands})
		}
	}

	return commands, nil
}
