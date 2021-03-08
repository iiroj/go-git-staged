package internal

import (
	"fmt"

	"github.com/bmatcuk/doublestar/v2"
)

// Command keeps a mapping of filenames and their commands
type Command struct {
	globs, files, commands []string
}

// CreateCommands resolves filenames to their commands
func CreateCommands(pairs []Pair, filenames []string) (commands []Command, err error) {
	commands = make([]Command, 0)

	for _, pair := range pairs {
		matches := make([]string, 0)
		for _, glob := range pair.globs {
			for _, filename := range filenames {
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

	for _, command := range commands {
		fmt.Println(command)
	}

	return commands, nil
}
