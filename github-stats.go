package main

import (
	"context"
	"fmt"
	"log"
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

	opt := &github.PullRequestListOptions{State: "closed", Sort: "created", Direction: "desc", ListOptions: github.ListOptions{PerPage: 100}}

	pulls, _, err := client.PullRequests.List(context.Background(), githubUsername, githubRepository, opt)

	if err != nil {
		log.Fatal(err)
	}

	var currentPull *github.PullRequest
	totalPullRequests := 0
	totalMergedInDays := 0.0

	for _, pull := range pulls {
		currentPull = pull
		createdAt := currentPull.GetCreatedAt()
		mergedAt := currentPull.GetMergedAt()
		mergedInDays := mergedAt.Sub(createdAt).Hours() / 24

		if mergedInDays > 0 {
			fmt.Println("Title:", currentPull.GetTitle())
			// fmt.Println("CreatedAt:", createdAt)
			// fmt.Println("MergedAt:", mergedAt)
			// fmt.Println(reflect.TypeOf(currentPull.GetMergedAt()))
			fmt.Printf("Merged in: %f days", mergedInDays)
			fmt.Printf("\n\n")
			totalMergedInDays += mergedInDays
			totalPullRequests++
		}
	}

	fmt.Printf("Total Pull Requests: %d \n", totalPullRequests)
	fmt.Printf("Average Time to Merge: %f days \n", totalMergedInDays/float64(totalPullRequests))

	// data, _ := json.MarshalIndent(repos, "", "  ")
	// fmt.Println(data)
	// fmt.Println("Errors: ", err)
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Printf("\n\n")

	return command
}
