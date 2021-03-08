package main

import (
	"os"

	"github.com/iiroj/go-git-staged/cmd"
)

func main() {
	if error := cmd.Execute(os.Args[1:]); error != nil {
		os.Exit((1))
	}
}
