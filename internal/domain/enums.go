package domain

import "fmt"

type EventName int

const (
	EventNameUnknown EventName = iota
	EventNamePullRequest
	EventNameIssues
)

func ToEventName(s string) (EventName, error) {
	switch s {
	case "pull_request":
		return EventNamePullRequest, nil
	case "issues":
		return EventNameIssues, nil
	default:
		return EventNameUnknown, fmt.Errorf("invalid event name: %s", s)
	}
}

func (e EventName) String() string {
	switch e {
	case EventNamePullRequest:
		return "pull_request"
	case EventNameIssues:
		return "issues"
	default:
		return "unknown"
	}
}

type Locale int

const (
	LocaleUnknown Locale = iota
	LocaleJa
	LocaleEn
)

func ToLocale(s string) (Locale, error) {
	switch s {
	case "ja":
		return LocaleJa, nil
	case "en":
		return LocaleEn, nil
	default:
		return LocaleUnknown, fmt.Errorf("invalid locale: %s", s)
	}
}

func (l Locale) String() string {
	switch l {
	case LocaleJa:
		return "ja"
	case LocaleEn:
		return "en"
	default:
		return "unknown"
	}
}
