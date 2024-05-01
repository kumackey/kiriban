package domain_test

import (
	"context"
	"testing"

	"github.com/kumackey/kiriban/internal/domain"
	mock "github.com/kumackey/kiriban/internal/domain/mock"
	"github.com/kumackey/kiriban/kiriban"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const (
	expectedJa = `おめでとうございます！🎉 #1000 はキリ番です！
次のキリ番は #1111 です。踏み逃げは厳禁ですよ！

| キリ番 | アカウント |
| --- | --- |
| #777 | [user1](https://example.com/user1) |
| #789 | [user2](https://example.com/user2) |
| #800 | [user3](https://example.com/user3) |
| #876 | [user4](https://example.com/user4) |
| #888 | [user5](https://example.com/user5) |
| #900 | [user6](https://example.com/user6) |
| #987 | [user7](https://example.com/user7) |
| #999 | [user1](https://example.com/user1) |
| #1000 | [user2](https://example.com/user2) |
| #1111 | まもなく…… |
`
	expectedEn = `Congratulations!🎉 #1000 is kiriban!
Next kiriban is #1111. But fleeing after stepping on kiriban is strictly forbidden, you know!

| kiriban | account |
| --- | --- |
| #777 | [user1](https://example.com/user1) |
| #789 | [user2](https://example.com/user2) |
| #800 | [user3](https://example.com/user3) |
| #876 | [user4](https://example.com/user4) |
| #888 | [user5](https://example.com/user5) |
| #900 | [user6](https://example.com/user6) |
| #987 | [user7](https://example.com/user7) |
| #999 | [user1](https://example.com/user1) |
| #1000 | [user2](https://example.com/user2) |
| #1111 | Comming Soon... |
`
)

func TestIssueCommenter_Message(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock.NewMockGitHubClient(ctrl)

	checker, err := kiriban.NewChecker()
	assert.NoError(t, err)
	ic := domain.NewIssueCommenter(mockClient, checker)

	v := 1000
	repo, err := domain.NewRepository("kumackey/example")
	assert.NoError(t, err)

	u := func(username string) domain.User {
		return domain.NewUser(username, "https://example.com/"+username)
	}

	users := map[int]domain.User{
		777:  u("user1"),
		789:  u("user2"),
		800:  u("user3"),
		876:  u("user4"),
		888:  u("user5"),
		900:  u("user6"),
		987:  u("user7"),
		999:  u("user1"),
		1000: u("user2"),
	}
	mockClient.EXPECT().GetIssueUsers(context.TODO(), repo, gomock.Any()).Return(users, nil).AnyTimes()

	testCases := []struct {
		locale          domain.Locale
		expectedComment string
	}{
		{
			locale:          domain.LocaleJa,
			expectedComment: expectedJa,
		},
		{
			locale:          domain.LocaleEn,
			expectedComment: expectedEn,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.locale.String(), func(t *testing.T) {
			comment, err := ic.Message(context.TODO(), repo, v, tc.locale)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedComment, comment)
		})
	}
}
