package char

import "github.com/fatih/color"

var (
	blue    = color.New(color.FgBlue).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	red     = color.New(color.FgRed).SprintFunc()

	// DoneChar is a green ✔︎
	Done = green("✔︎")
	// FailChar is a red ×
	Fail = red("×")
	// InfoChar is a blue ℹ︎✔
	Info = blue("ℹ︎")
	// RunChar is a magenta ❯
	Run = magenta("❯")
)
