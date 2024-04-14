package main

import "fmt"

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
	case "issues":
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
		return "issues"
	default:
		return "unknown"
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
