package database

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

type mockConn struct {
	closeCalled bool
}

func (m *mockConn) Close(ctx context.Context) error {
	m.closeCalled = true
	return nil
}

func TestUserNameExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockDB(ctrl)

	expectedUser := "testuser"

	mockDB.EXPECT().
		Connect().
		Return(mockDB, nil)

	mockDB.EXPECT().
		verifyUsername(mockDB, expectedUser).
		Return(true)

	got := UserNameExists(expectedUser)
	if got != true {
		t.Errorf("UserNameExists(%q) = %v; want true", expectedUser, got)
	}
}
