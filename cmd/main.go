package main

import (
	"context"
	"flag"
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
	e := flag.String("e", "", "Event name")
	flag.Parse()

	en, err := toEventName(*e)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Event name: %s\n", en.String())

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("Invalid arguments")
	}

	issueNumber, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	d, err := kiriban.NewDeterminator(kiriban.EnableDigitBasedRoundDetermination())
	if err != nil {
		log.Fatalln(err)
	}

	if !d.IsKiriban(issueNumber) {
		fmt.Printf("#%d is not kiriban.\n", issueNumber)
		os.Exit(0)
	}

	fmt.Printf("#%d is kiriban!\n", issueNumber)

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
		fmt.Sprintf("Congratulations! #%d is kiriban! 🎉\nNext kiriban is #%d.\n", issueNumber, d.Next(issueNumber)),
	)}

	ic, _, err := client.Issues.CreateComment(ctx, owner, repo, issueNumber, comment)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Commented: %s\n", *ic.HTMLURL)
}

type eventName int

const (
	unknown eventName = iota
	pullRequest
	issues
)

func toEventName(s string) (eventName, error) {
	switch s {
	case "pull_request":
		return pullRequest, nil
	case "issues":
		return issues, nil
	default:
		return unknown, fmt.Errorf("invalid event name: %s", s)
	}
}

func (e eventName) String() string {
	switch e {
	case pullRequest:
		return "pull_request"
	case issues:
		return "issues"
	default:
		return "unknown"
	}
}
