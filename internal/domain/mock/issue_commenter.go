// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/domain/issue_commenter.go
//
// Generated by this command:
//
//	mockgen -source=../internal/domain/issue_commenter.go -destination=../internal/domain/mock/issue_commenter.go -package=domainmock
//

// Package domainmock is a generated GoMock package.
package domainmock

import (
	context "context"
	reflect "reflect"

	domain "github.com/kumackey/kiriban/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockGitHubClient is a mock of GitHubClient interface.
type MockGitHubClient struct {
	ctrl     *gomock.Controller
	recorder *MockGitHubClientMockRecorder
}

// MockGitHubClientMockRecorder is the mock recorder for MockGitHubClient.
type MockGitHubClientMockRecorder struct {
	mock *MockGitHubClient
}

// NewMockGitHubClient creates a new mock instance.
func NewMockGitHubClient(ctrl *gomock.Controller) *MockGitHubClient {
	mock := &MockGitHubClient{ctrl: ctrl}
	mock.recorder = &MockGitHubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitHubClient) EXPECT() *MockGitHubClientMockRecorder {
	return m.recorder
}

// CreateIssueComment mocks base method.
func (m *MockGitHubClient) CreateIssueComment(arg0 context.Context, arg1 domain.Repository, arg2 int, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIssueComment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIssueComment indicates an expected call of CreateIssueComment.
func (mr *MockGitHubClientMockRecorder) CreateIssueComment(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIssueComment", reflect.TypeOf((*MockGitHubClient)(nil).CreateIssueComment), arg0, arg1, arg2, arg3)
}

// GetIssueUsers mocks base method.
func (m *MockGitHubClient) GetIssueUsers(arg0 context.Context, arg1 domain.Repository, arg2 []int) (map[int]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueUsers", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[int]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssueUsers indicates an expected call of GetIssueUsers.
func (mr *MockGitHubClientMockRecorder) GetIssueUsers(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueUsers", reflect.TypeOf((*MockGitHubClient)(nil).GetIssueUsers), arg0, arg1, arg2)
}
