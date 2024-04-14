package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/kumackey/kiriban/kiriban"
)

type IssueCommenter struct {
	client GitHubClient
	kc     *kiriban.Checker
}

type GitHubClient interface {
	CreateIssueComment(context.Context, repository, int, string) (string, error)
	GetIssueUsers(context.Context, repository, []int) (map[int]string, error)
}

func NewIssueCommenter(client GitHubClient, kc *kiriban.Checker) IssueCommenter {
	return IssueCommenter{client: client, kc: kc}
}

// TODO: test
func (ic IssueCommenter) Execute(ctx context.Context, cfg config, v int) (string, error) {
	msg, err := ic.message(ctx, cfg.repository, v, cfg.locale)
	if err != nil {
		return "", err
	}

	return ic.client.CreateIssueComment(ctx, cfg.repository, v, msg)
}

func (ic IssueCommenter) message(ctx context.Context, repository repository, v int, l locale) (string, error) {
	var msg string
	next := ic.kc.Next(v)

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

	users, err := ic.client.GetIssueUsers(ctx, repository, list)
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

func (ic IssueCommenter) calcPreviousKiribans(number, limit int) []int {
	list := make([]int, 0, limit+2) // +2 is for the current kiriban and the next kiriban

	for limit > 0 {
		num, err := ic.kc.Previous(number)
		if errors.Is(err, kiriban.ErrorNoPreviousKiriban) {
			break
		}
		list = append([]int{num}, list...)
		number = num
		limit--
	}

	return list
}
