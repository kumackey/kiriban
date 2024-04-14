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
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	t := flag.String("t", "", "Event name")
	udks := flag.String("u", "", "User-defined kiribans")
	flag.Parse()

	_, err = toEventName(*t)
	if err != nil {
		log.Fatalln(err)
	}

	eks, err := toExceptionalKiribans(*udks)
	if err != nil {
		log.Fatalln(err)
	}

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("Invalid arguments")
	}

	issueNumber, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	c, err := kiriban.NewChecker(kiriban.EnableDigitBasedRoundDetermination(), kiriban.ExceptionalKiribanOption(eks))
	if err != nil {
		log.Fatalln(err)
	}

	if !c.IsKiriban(issueNumber) {
		fmt.Printf("#%c is not kiriban.\n", issueNumber)
		os.Exit(0)
	}

	fmt.Printf("#%c is kiriban!\n", issueNumber)

	ctx := context.Background()

	ic := NewIssueCommenter(newGithubClient(ctx, cfg.githubToken), c)
	url, err := ic.Execute(ctx, cfg, issueNumber)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Commented: %s\n", url)
}

func toExceptionalKiribans(e string) ([]kiriban.ExceptionalKiriban, error) {
	if e == "" {
		return nil, nil
	}

	e = strings.ReplaceAll(e, " ", "")
	parts := strings.Split(e, ",")
	kiribans := make([]kiriban.ExceptionalKiriban, 0, len(parts))

	for _, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid exceptional kiriban value: %v", err)
		}

		kiribans = append(kiribans, kiriban.ExceptionalKiriban{Value: value})
	}

	return kiribans, nil
}

type githubClientImpl struct {
	client *github.Client
}

func newGithubClient(ctx context.Context, githubToken string) *githubClientImpl {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &githubClientImpl{client: client}
}

func (g *githubClientImpl) CreateIssueComment(ctx context.Context, repository repository, number int, comment string) (string, error) {
	c := &github.IssueComment{Body: &comment}
	issueComment, _, err := g.client.Issues.CreateComment(ctx, repository.owner, repository.repo, number, c)
	if err != nil {
		return "", err
	}

	return issueComment.GetHTMLURL(), nil
}

func (g *githubClientImpl) GetIssueUsers(ctx context.Context, repository repository, numbers []int) (map[int]string, error) {
	users := make(map[int]string, len(numbers))
	for _, number := range numbers {
		// TODO: N+1 problem
		issue, _, err := g.client.Issues.Get(ctx, repository.owner, repository.repo, number)
		if err != nil {
			return nil, err
		}

		users[number] = issue.GetUser().GetLogin()
	}

	return users, nil
}
