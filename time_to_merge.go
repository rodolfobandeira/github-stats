package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func timeToMerge(githubToken string, githubUsername string, githubRepository string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.PullRequestListOptions{
		State:     "closed",
		Sort:      "merged",
		Direction: "desc",
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	pulls, _, err := client.PullRequests.List(
		context.Background(),
		githubUsername,
		githubRepository,
		opt,
	)

	if err != nil {
		log.Fatal(err)
	}

	var currentPull *github.PullRequest
	totalPullRequests := 0
	totalMergedInDays := 0.0

	file, err := os.OpenFile(
		"stats_reports/time_to_merge.csv",
		os.O_CREATE|os.O_RDWR,
		0755,
	)

	csvHeader := fmt.Sprintf(
		"\x22%s\x22,\x22%s\x22\n",
		"PR Title",
		"Merged in Days",
	)

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
			csvRow := fmt.Sprintf(
				"\x22%s\x22,\x22%f\x22\n",
				strings.Trim(pullRequestTitle, " "),
				mergedInDays,
			)

			file.WriteString(csvRow)
			fmt.Printf(csvRow)
			// fmt.Println("Title:", pullRequestTitle)
			// fmt.Println("CreatedAt:", createdAt)
			// fmt.Println("MergedAt:", mergedAt)
			// fmt.Println(reflect.TypeOf(currentPull.GetMergedAt()))
			// fmt.Printf("Merged in: %f days", mergedInDays)
			// fmt.Printf("\n\n")
			totalMergedInDays += mergedInDays
			totalPullRequests++
		}
	}

	file.Close()

	emptyLine()
	c := color.New(color.FgYellow).Add(color.BgBlack).Add(color.Bold)
	fmt.Printf("Total Pull Requests: %d \n", totalPullRequests)
	c.Printf("Average Time to Merge: %f days", totalMergedInDays/float64(totalPullRequests))
	emptyLine()

	// data, _ := json.MarshalIndent(repos, "", "  ")
	// fmt.Println(data)
	// fmt.Println("Errors: ", err)
}
