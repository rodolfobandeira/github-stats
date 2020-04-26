package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rodolfobandeira/github-stats/stats"
	"github.com/rodolfobandeira/github-stats/utils"
)

func main() {
	utils.IntroLogo()

	githubTokenFlag := flag.String("t", os.Getenv("GITHUB_TOKEN"), "-t <token> Your Github token (Required)")
	githubUserFlag := flag.String("u", os.Getenv("GITHUB_USER"), "-u <username> Your Github username (Required)")
	githubRepoFlag := flag.String("r", os.Getenv("GITHUB_REPO"), "-r <repository> Repository you want to get stats (Required)")

	flag.Parse()

	if utils.ContainsEmpty(*githubTokenFlag, *githubUserFlag, *githubRepoFlag) {
		showOptions()
		os.Exit(0)
	}

	stats.TimeToMerge(*githubTokenFlag, *githubUserFlag, *githubRepoFlag)

	utils.EmptyLine()
	os.Exit(0)
}

func showOptions() {
	c := color.New(color.FgGreen).Add(color.BgBlack).Add(color.Bold)
	c.Printf("github-stats -u username -r repository -s ttm -t github_token")
	utils.EmptyLine()

	utils.EmptyLine()
	fmt.Println("-u <username> or -o <organization> (Required)")
	fmt.Println("-r <repository> Here you should put the repository name. (Required)")
	fmt.Println("-s Which stats report do you want to see?")
	fmt.Println("	<ttm> Time to Merge Pull Requests in days.")
	fmt.Println("-t Github Token")
	utils.EmptyLine()
}
