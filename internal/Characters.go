package internal

import "github.com/fatih/color"

var (
	magenta = color.New(color.FgMagenta).SprintFunc()
	// RunChar is a magenta ❯
	RunChar = magenta("❯")
	red     = color.New(color.FgRed).SprintFunc()
	// FailChar is a red ×
	FailChar = red("×")
	green    = color.New(color.FgGreen).SprintFunc()
	// DoneChar is a green ✔︎
	DoneChar = green("✔︎")
	blue     = color.New(color.FgBlue).SprintFunc()
	// InfoChar is a blue ℹ︎✔
	InfoChar = blue("ℹ︎")
)
