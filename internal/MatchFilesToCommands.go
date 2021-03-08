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

	for _, pair := range pairs {
		matches := make([]string, 0)
		for _, glob := range pair.globs {
			for _, filename := range files {
				match, matchError := doublestar.PathMatch(glob, filename)
				if matchError != nil {
					return nil, matchError
				}
				if match == true {
					matches = append(matches, filename)
				}
			}
		}

		if len(matches) > 0 {
			commands = append(commands, Command{pair.globs, matches, pair.commands})
		}
	}

	return commands, nil
}
