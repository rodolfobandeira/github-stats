package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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

	opt := &github.PullRequestListOptions{State: "closed", Sort: "merged", Direction: "desc", ListOptions: github.ListOptions{PerPage: 100}}

	pulls, _, err := client.PullRequests.List(context.Background(), githubUsername, githubRepository, opt)

	if err != nil {
		log.Fatal(err)
	}

	var currentPull *github.PullRequest
	totalPullRequests := 0
	totalMergedInDays := 0.0

	file, err := os.OpenFile("stats_reports/time_to_merge.csv", os.O_CREATE|os.O_RDWR, 0755)
	csvHeader := fmt.Sprintf("\x22%s\x22,\x22%s\x22\n", "PR Title", "Merged in Days")

	file.WriteString(csvHeader)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	for _, pull := range pulls {
		currentPull = pull
		createdAt := currentPull.GetCreatedAt()
		mergedAt := currentPull.GetMergedAt()
		mergedInDays := mergedAt.Sub(createdAt).Hours() / 24

		if mergedInDays > 0 {
			pullRequestTitle := sanitizeLine(currentPull.GetTitle())
			csvRow := fmt.Sprintf("\x22%s\x22,\x22%f\x22\n", pullRequestTitle, mergedInDays)

			file.WriteString(csvRow)
			fmt.Println("Title:", pullRequestTitle)
			// fmt.Println("CreatedAt:", createdAt)
			// fmt.Println("MergedAt:", mergedAt)
			// fmt.Println(reflect.TypeOf(currentPull.GetMergedAt()))
			fmt.Printf("Merged in: %f days", mergedInDays)
			fmt.Printf("\n\n")
			totalMergedInDays += mergedInDays
			totalPullRequests++
		}
	}

	file.Close()

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

func sanitizeLine(lineRow string) string {
	lineRow = strings.ReplaceAll(lineRow, "\"", "'")
	lineRow = strings.ReplaceAll(lineRow, ",", ".")

	return lineRow
}
