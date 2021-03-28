package main

import (
	"os"

	"github.com/iiroj/go-git-staged/cmd"
)

func main() {
	if errors := cmd.Execute(os.Args[1:]); errors > 0 {
		os.Exit((1))
	}
}
