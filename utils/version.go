package utils

import (
	"fmt"

	"github.com/fatih/color"
)

const currentVersion float32 = 0.1 // This is how to do constants in Go.

// ShowVersion - Shows current version of the application
func ShowVersion() {
	IntroLogo()
	fmt.Printf("Version: %f\n", currentVersion)
	fmt.Println("https://github.com/rodolfobandeira/github-stats")
}

func IntroLogo() {
	c := color.New(color.FgRed).Add(color.BgBlack).Add(color.Bold)
	c.Print("╔═╗┬┌┬┐┬ ┬┬ ┬┌┐   ╔═╗┌┬┐┌─┐┌┬┐┌─┐")
	EmptyLine()
	c.Print("║ ╦│ │ ├─┤│ │├┴┐  ╚═╗ │ ├─┤ │ └─┐")
	EmptyLine()
	c.Print("╚═╝┴ ┴ ┴ ┴└─┘└─┘  ╚═╝ ┴ ┴ ┴ ┴ └─┘")
	EmptyLine()
}
