package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type repository struct {
	owner string
	repo  string
}

type config struct {
	githubToken string
	repository  repository
	locale      locale
}

func loadConfig() (config, error) {
	token, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		return config{}, errors.New("missing GITHUB_TOKEN")
	}

	repo, ok := os.LookupEnv("GITHUB_REPOSITORY")
	if !ok {
		return config{}, errors.New("missing GITHUB_REPOSITORY")
	}

	parts := strings.Split(repo, "/")
	if len(parts) != 2 {
		return config{}, fmt.Errorf("invalid GITHUB_REPOSITORY: %s", repo)
	}

	loc, ok := os.LookupEnv("LOCALE")
	if !ok {
		loc = localeEn.String()
	}
	lcl, err := toLocale(loc)
	if err != nil {
		return config{}, err
	}

	return config{
		githubToken: token,
		repository:  repository{owner: parts[0], repo: parts[1]},
		locale:      lcl,
	}, nil
}
