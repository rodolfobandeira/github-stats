package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	introLogo()

	githubTokenFlag := flag.String("t", os.Getenv("GITHUB_TOKEN"), "-t <token> Your Github token (Required)")
	githubUserFlag := flag.String("u", os.Getenv("GITHUB_USER"), "-u <username> Your Github username (Required)")
	githubRepoFlag := flag.String("r", os.Getenv("GITHUB_REPO"), "-r <repository> Repository you want to get stats (Required)")

	flag.Parse()

	if containsEmpty(*githubTokenFlag, *githubUserFlag, *githubRepoFlag) {
		showOptions()
		os.Exit(0)
	}

	timeToMerge(*githubTokenFlag, *githubUserFlag, *githubRepoFlag)

	emptyLine()
	os.Exit(0)
}

func introLogo() {
	c := color.New(color.FgRed).Add(color.BgBlack).Add(color.Bold)
	c.Print("╔═╗┬┌┬┐┬ ┬┬ ┬┌┐   ╔═╗┌┬┐┌─┐┌┬┐┌─┐")
	emptyLine()
	c.Print("║ ╦│ │ ├─┤│ │├┴┐  ╚═╗ │ ├─┤ │ └─┐")
	emptyLine()
	c.Print("╚═╝┴ ┴ ┴ ┴└─┘└─┘  ╚═╝ ┴ ┴ ┴ ┴ └─┘")
	emptyLine()
}

func showOptions() {
	c := color.New(color.FgGreen).Add(color.BgBlack).Add(color.Bold)
	c.Printf("github-stats -u username -r repository -s ttmd -o ttmd.csv -t github_token")
	emptyLine()

	emptyLine()
	fmt.Println("-u <username> or -o <organization> (Required)")
	fmt.Println("-r <repository> Here you should put the repository name. (Required)")
	fmt.Println("-s Which stats report do you want to see?")
	fmt.Println("	<ttmd> Time to Merge Pull Requests in days.")
	fmt.Println("	<ttmh> Time to Merge Pull Requests in hours.")
	fmt.Println("-o <filename.csv> Write the output to file using CSV format.")
	fmt.Println("-h Help me!")
	fmt.Println("-v Shows github-stats version.")
	emptyLine()
}
