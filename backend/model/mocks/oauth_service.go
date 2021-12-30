package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockOAuthService struct {
	mock.Mock
}

func (m *MockOAuthService) GetLoginURL(ctx context.Context, provider string, state string) (string, error) {
	ret := m.Called(ctx, provider, state)

	var r0 string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1

}
