package main

import (
	"os"

	"github.com/iiroj/go-git-staged/cmd"
)

// Run go-git-staged cli
func main() {
	// Main cmd contains error handling, and a spinner
	if result := cmd.Execute(os.Args[1:]); result.Error != nil {
		os.Exit((1))
	}
}
