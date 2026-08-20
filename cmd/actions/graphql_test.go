package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kumackey/kiriban/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildIssueUsersQuery(t *testing.T) {
	expected := `query($owner: String!, $repo: String!) {
  repository(owner: $owner, name: $repo) {
    n10: issueOrPullRequest(number: 10) {
      ... on Issue { author { login url } }
      ... on PullRequest { author { login url } }
    }
    n20: issueOrPullRequest(number: 20) {
      ... on Issue { author { login url } }
      ... on PullRequest { author { login url } }
    }
  }
}`

	assert.Equal(t, expected, buildIssueUsersQuery([]int{10, 20}))
}

func TestGetIssueUsers(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		numbers  []int
		response string
		expected map[int]domain.User
		wantErr  bool
	}{
		"issues and pull requests are fetched in one request": {
			numbers: []int{10, 20},
			response: `{"data":{"repository":{
				"n10":{"author":{"login":"user1","url":"https://example.com/user1"}},
				"n20":{"author":{"login":"user2","url":"https://example.com/user2"}}
			}}}`,
			expected: map[int]domain.User{
				10: domain.NewUser("user1", "https://example.com/user1"),
				20: domain.NewUser("user2", "https://example.com/user2"),
			},
		},
		"a missing number is skipped instead of failing the whole request": {
			numbers: []int{10, 20},
			response: `{"data":{"repository":{
				"n10":{"author":{"login":"user1","url":"https://example.com/user1"}},
				"n20":null
			}},"errors":[{"type":"NOT_FOUND","message":"Could not resolve to an issue or pull request with the number of 20."}]}`,
			expected: map[int]domain.User{
				10: domain.NewUser("user1", "https://example.com/user1"),
			},
		},
		"a deleted author is skipped": {
			numbers:  []int{10},
			response: `{"data":{"repository":{"n10":{"author":null}}}}`,
			expected: map[int]domain.User{},
		},
		"an error other than NOT_FOUND fails": {
			numbers:  []int{10},
			response: `{"data":{"repository":null},"errors":[{"type":"FORBIDDEN","message":"Resource not accessible"}]}`,
			wantErr:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var requests int
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests++

				body, err := io.ReadAll(r.Body)
				require.NoError(t, err)

				var req graphQLRequest
				require.NoError(t, json.Unmarshal(body, &req))
				assert.Equal(t, "kumackey", req.Variables["owner"])
				assert.Equal(t, "kiriban", req.Variables["repo"])

				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(tt.response))
			}))
			defer srv.Close()

			c := &githubClientImpl{httpClient: srv.Client(), graphQLEndpoint: srv.URL}
			users, err := c.GetIssueUsers(context.Background(), domain.Repository{Owner: "kumackey", Repo: "kiriban"}, tt.numbers)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, users)
			assert.Equal(t, 1, requests, "all the numbers must be fetched in a single request")
		})
	}
}
