package main

import "fmt"

const currentVersion float32 = 0.1 // This is how to do constants in Go.

func showVersion() {
	introLogo()
	fmt.Printf("Version: %f\n", currentVersion)
	fmt.Println("https://github.com/rodolfobandeira/github-stats")
}
