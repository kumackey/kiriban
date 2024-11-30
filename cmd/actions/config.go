package main

import (
	"errors"
	"github.com/kumackey/kiriban/internal/domain"
	"os"
)

type config struct {
	githubToken string
	repository  domain.Repository
	locale      domain.Locale
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

	rp, err := domain.NewRepository(repo)
	if err != nil {
		return config{}, err
	}

	loc, ok := os.LookupEnv("LOCALE")
	if !ok {
		loc = domain.LocaleEn.String()
	}
	lcl, err := domain.ToLocale(loc)
	if err != nil {
		return config{}, err
	}

	return config{githubToken: token, repository: rp, locale: lcl}, nil
}
