package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/kumackey/kiriban/kiriban"
	"golang.org/x/oauth2"
	"log"
)

type issueCommenter struct {
	client *github.Client
	kd     *kiriban.Determinator
}

func newIssueCommenter(client *github.Client, kd *kiriban.Determinator) issueCommenter {
	return issueCommenter{client: client, kd: kd}
}

func newGithubClient(ctx context.Context, githubToken string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return client
}

func (ic issueCommenter) execute(ctx context.Context, cfg config, v int) (*github.IssueComment, error) {
	msg, err := ic.message(ctx, cfg.repository, v, cfg.locale)
	if err != nil {
		return nil, err
	}

	comment := &github.IssueComment{Body: github.String(msg)}
	c, _, err := ic.client.Issues.CreateComment(ctx, cfg.repository.owner, cfg.repository.repo, v, comment)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// TODO: required to test and refactor
func (ic issueCommenter) message(ctx context.Context, repository repository, v int, l locale) (string, error) {
	var msg string
	next := ic.kd.Next(v)

	switch l {
	case localeJa:
		msg = fmt.Sprintf("おめでとうございます！🎉 #%d はキリ番です！\n次のキリ番は #%d です。踏み逃げは厳禁ですよ！\n", v, next)
	case localeEn:
		msg = fmt.Sprintf("Congratulations!🎉 #%d is kiriban!\nNext kiriban is #%d. But fleeing after stepping on kiriban is strictly forbidden, you know!\n", v, next)
	default:
		return "", fmt.Errorf("unsupported locale: %s", l.String())
	}

	list := ic.calcPreviousKiribans(v, 8)
	list = append(list, v)

	users, err := ic.fetchIssueUsers(ctx, repository, list)
	if err != nil {
		log.Fatalln(err)
	}

	switch l {
	case localeJa:
		msg += "\n| キリ番 | アカウント |\n| --- | --- |\n"
	case localeEn:
		msg += "\n| kiriban | account |\n| --- | --- |\n"
	default:
		return "", fmt.Errorf("unsupported locale: %s", l.String())
	}

	for _, l := range list {
		msg += fmt.Sprintf("| #%d | @%s |\n", l, users[l])
	}

	switch l {
	case localeJa:
		msg += fmt.Sprintf("| #%d | まもなく…… |\n", next)
	case localeEn:
		msg += fmt.Sprintf("| #%d | Comming Soon... |\n", next)
	default:
		return "", fmt.Errorf("unsupported locale: %s", l.String())
	}

	return msg, nil
}

func (ic issueCommenter) calcPreviousKiribans(number, limit int) []int {
	list := make([]int, 0, limit+2) // +2 is for the current kiriban and the next kiriban

	for limit > 0 {
		num, err := ic.kd.Previous(number)
		if errors.Is(err, kiriban.ErrorNoPreviousKiriban) {
			break
		}
		list = append([]int{num}, list...)
		number = num
		limit--
	}

	return list
}

func (ic issueCommenter) fetchIssueUsers(ctx context.Context, repository repository, numbers []int) (map[int]string, error) {
	users := make(map[int]string, len(numbers))

	for _, number := range numbers {
		// TODO: N+1 problem
		issue, _, err := ic.client.Issues.Get(ctx, repository.owner, repository.repo, number)
		if err != nil {
			// TODO: Handle error
			return nil, err
		}

		users[number] = issue.GetUser().GetLogin()
	}

	return users, nil
}
