package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/kumackey/kiriban/kiriban"
	"golang.org/x/oauth2"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	prNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	c, err := kiriban.NewChecker()
	if err != nil {
		log.Fatalln(err)
	}

	if !c.IsKiriban(prNumber) {
		fmt.Printf("#%d is not kiriban.\n", prNumber)
		os.Exit(0)
	}

	fmt.Printf("#%d is kiriban!\n", prNumber)

	ctx := context.Background()

	token, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		log.Fatalln("GITHUB_TOKEN is not set.")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repository, ok := os.LookupEnv("GITHUB_REPOSITORY")
	if !ok {
		log.Fatalln("GITHUB_REPOSITORY is not set.")
	}
	parts := strings.Split(repository, "/")
	if len(parts) != 2 {
		log.Fatalf("Invalid GITHUB_REPOSITORY: %s\n", repository)
	}
	owner, repo := parts[0], parts[1]

	comment := &github.IssueComment{Body: github.String(
		fmt.Sprintf("Congratulations! #%d is kiriban! 🎉", prNumber),
	)}

	ic, _, err := client.Issues.CreateComment(ctx, owner, repo, prNumber, comment)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Commented: %s\n", *ic.HTMLURL)
}
