package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	introMessage()

	for {
		showMenuOptions()
		command := readCommand()

		switch command {
		case 1:
			fmt.Println("TTMPR - Time to Merge Pull Requests")

		case 2:
			listPublicRepos()
		case 0:
			fmt.Println("Ok, bye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid Option")
			os.Exit(-1)
		}

		fmt.Printf("\n\n")
	}
}

func introMessage() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("Github Stats - Which statistics do you want to see?")
	fmt.Println("---------------------------------------------------")
}

func showMenuOptions() {
	fmt.Println("1- Time to Merge Pull Requests")
	fmt.Println("2- List Public Repositories")
	fmt.Println("0- Exit")
}

func listPublicRepos() {
	githubToken := os.Getenv("GITHUB_TOKEN")
	githubUsername := os.Getenv("GITHUB_USERNAME")
	githubRepository := os.Getenv("GITHUB_REPOSITORY")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "phpcanada", opt)

	client.PullRequests.
		fmt.Println("Repos: ", repos)
	fmt.Println("Errors: ", err)
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Printf("\n\n")

	return command
}
