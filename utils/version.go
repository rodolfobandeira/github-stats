package utils

import "fmt"

const currentVersion float32 = 0.1 // This is how to do constants in Go.

// ShowVersion - Shows current version of the application
func ShowVersion() {
	IntroLogo()
	fmt.Printf("Version: %f\n", currentVersion)
	fmt.Println("https://github.com/rodolfobandeira/github-stats")
}
