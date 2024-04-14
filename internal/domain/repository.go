package domain

import (
	"fmt"
	"strings"
)

type Repository struct {
	Owner string
	Repo  string
}

func NewRepository(repository string) (Repository, error) {
	parts := strings.Split(repository, "/")
	if len(parts) != 2 {
		return Repository{}, fmt.Errorf("invalid GITHUB_REPOSITORY: %s", repository)
	}

	return Repository{Owner: parts[0], Repo: parts[1]}, nil
}
