package main

import (
	"context"
	"errors"
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

	loc, ok := os.LookupEnv("LOCALE")
	if !ok {
		loc = localeEn.String()
	}
	lcl, err := toLocale(loc)
	if err != nil {
		log.Fatalln(err)
	}

	owner, repo := parts[0], parts[1]

	msg, err := message(ctx, client, d, owner, repo, issueNumber, lcl)
	if err != nil {
		log.Fatalln(err)
	}
	comment := &github.IssueComment{Body: github.String(msg)}

	ic, _, err := client.Issues.CreateComment(ctx, owner, repo, issueNumber, comment)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Commented: %s\n", *ic.HTMLURL)
}

// TODO: required to test and refactor
func message(ctx context.Context, client *github.Client, d *kiriban.Determinator, owner, repo string, v int, l locale) (string, error) {
	var msg string
	next := d.Next(v)

	switch l {
	case localeJa:
		msg = fmt.Sprintf("おめでとうございます！🎉 #%d はキリ番です！\n次のキリ番は #%d です。踏み逃げは厳禁ですよ！\n", v, next)
	case localeEn:
		msg = fmt.Sprintf("Congratulations!🎉 #%d is kiriban!\nNext kiriban is #%d. But fleeing after stepping on kiriban is strictly forbidden, you know!\n", v, next)
	default:
		return "", fmt.Errorf("unsupported locale: %s", l.String())
	}

	list := calcPreviousKiribans(d, v, 8)
	list = append(list, v)

	users, err := fetchIssueUsers(ctx, client, owner, repo, list)
	if err != nil {
		log.Fatalln(err)
	}

	msg += "\n| kiriban | account |\n| --- | --- |\n"

	for _, l := range list {
		msg += fmt.Sprintf("| #%d | @%s |\n", l, users[l])
	}

	msg += fmt.Sprintf("| #%d | Comming Soon... |\n", next)

	return msg, nil
}

func calcPreviousKiribans(d *kiriban.Determinator, number, limit int) []int {
	list := make([]int, 0, limit+2) // +2 is for the current kiriban and the next kiriban

	for limit > 0 {
		num, err := d.Previous(number)
		if errors.Is(err, kiriban.ErrorNoPreviousKiriban) {
			break
		}
		list = append([]int{num}, list...)
		limit--

	}

	return list
}

func fetchIssueUsers(ctx context.Context, client *github.Client, owner, repo string, numbers []int) ([]string, error) {
	users := make([]string, 0, len(numbers))

	for _, number := range numbers {
		issue, _, err := client.Issues.Get(ctx, owner, repo, number)
		if err != nil {
			// TODO: Handle error
			return nil, err
		}

		users = append(users, issue.User.GetLogin())
	}

	return users, nil
}

type eventName int

const (
	eventNameUnknown eventName = iota
	eventNamePullRequest
	eventNameIssues
)

func toEventName(s string) (eventName, error) {
	switch s {
	case "pull_request":
		return eventNamePullRequest, nil
	case "eventNameIssues":
		return eventNameIssues, nil
	default:
		return eventNameUnknown, fmt.Errorf("invalid event name: %s", s)
	}
}

func (e eventName) String() string {
	switch e {
	case eventNamePullRequest:
		return "pull_request"
	case eventNameIssues:
		return "eventNameIssues"
	default:
		return "eventNameUnknown"
	}
}

type locale int

const (
	localeUnknown locale = iota
	localeJa
	localeEn
)

func toLocale(s string) (locale, error) {
	switch s {
	case "ja":
		return localeJa, nil
	case "en":
		return localeEn, nil
	default:
		return localeUnknown, fmt.Errorf("invalid locale: %s", s)
	}
}

func (l locale) String() string {
	switch l {
	case localeJa:
		return "ja"
	case localeEn:
		return "en"
	default:
		return "unknown"
	}
}
