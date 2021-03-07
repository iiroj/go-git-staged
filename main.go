package main

import (
	"log"
	"os"

	"github.com/iiroj/go-git-staged/cmd"
)

// Run go-git-staged cli
func main() {
	// Main cmd contains error handling, and a spinner
	result := cmd.Execute(os.Args[1:])

	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
