package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kumackey/kiriban/internal/domain"
)

const githubGraphQLEndpoint = "https://api.github.com/graphql"

// notFoundErrorType is the error type GitHub returns for an alias that points to
// a number which does not exist. It arrives alongside the data of the other
// aliases, so it is tolerated instead of failing the whole request.
const notFoundErrorType = "NOT_FOUND"

type graphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
}

type graphQLError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type issueUsersResponse struct {
	Data struct {
		Repository map[string]*struct {
			Author *struct {
				Login string `json:"login"`
				URL   string `json:"url"`
			} `json:"author"`
		} `json:"repository"`
	} `json:"data"`
	Errors []graphQLError `json:"errors"`
}

// GetIssueUsers fetches the authors of the given issues and pull requests.
// All the numbers are looked up in a single GraphQL request by aliasing each
// lookup, so the number of API calls does not grow with the number of kiribans.
func (g *githubClientImpl) GetIssueUsers(ctx context.Context, repository domain.Repository, numbers []int) (map[int]domain.User, error) {
	users := make(map[int]domain.User, len(numbers))
	if len(numbers) == 0 {
		return users, nil
	}

	numbers = uniqueNumbers(numbers)

	req := graphQLRequest{
		Query: buildIssueUsersQuery(numbers),
		Variables: map[string]any{
			"owner": repository.Owner,
			"repo":  repository.Repo,
		},
	}

	var res issueUsersResponse
	if err := g.doGraphQL(ctx, req, &res); err != nil {
		return nil, err
	}

	// An alias may be null when its number does not exist, which GitHub reports
	// as a NOT_FOUND error next to the data of the other aliases. Any other
	// error means the response cannot be trusted.
	for _, e := range res.Errors {
		if e.Type != notFoundErrorType {
			return nil, fmt.Errorf("github graphql error: %s", e.Message)
		}
	}

	if res.Data.Repository == nil {
		return nil, fmt.Errorf("repository %s/%s not found", repository.Owner, repository.Repo)
	}

	for _, number := range numbers {
		node := res.Data.Repository[issueUserAlias(number)]
		if node == nil || node.Author == nil {
			continue
		}

		users[number] = domain.NewUser(node.Author.Login, node.Author.URL)
	}

	return users, nil
}

func (g *githubClientImpl) doGraphQL(ctx context.Context, req graphQLRequest, res any) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, g.graphQLEndpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpRes, err := g.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()

	// A GraphQL error is reported in the body with a 200, so a non-200 status
	// means the request itself was rejected.
	if httpRes.StatusCode != http.StatusOK {
		return fmt.Errorf("github graphql request failed: %s", httpRes.Status)
	}

	return json.NewDecoder(httpRes.Body).Decode(res)
}

// buildIssueUsersQuery builds a query that looks up every number at once.
// issueOrPullRequest is used because a kiriban number may be either of them.
func buildIssueUsersQuery(numbers []int) string {
	var b strings.Builder

	b.WriteString("query($owner: String!, $repo: String!) {\n")
	b.WriteString("  repository(owner: $owner, name: $repo) {\n")

	for _, number := range numbers {
		fmt.Fprintf(&b, "    %s: issueOrPullRequest(number: %d) {\n", issueUserAlias(number), number)
		b.WriteString("      ... on Issue { author { login url } }\n")
		b.WriteString("      ... on PullRequest { author { login url } }\n")
		b.WriteString("    }\n")
	}

	b.WriteString("  }\n")
	b.WriteString("}")

	return b.String()
}

// issueUserAlias converts a number into a GraphQL alias, which must not start
// with a digit.
func issueUserAlias(number int) string {
	return fmt.Sprintf("n%d", number)
}

func uniqueNumbers(numbers []int) []int {
	seen := make(map[int]struct{}, len(numbers))
	uniques := make([]int, 0, len(numbers))

	for _, number := range numbers {
		if _, ok := seen[number]; ok {
			continue
		}

		seen[number] = struct{}{}
		uniques = append(uniques, number)
	}

	return uniques
}
