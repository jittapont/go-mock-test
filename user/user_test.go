package user_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jittapont/go-mock-test/mocks"
	"github.com/jittapont/go-mock-test/user"
)

func Test_insertUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepository := mocks.NewMockUserRepository(ctrl)
	u := &user.User{
		Name: "test",
	}
	mockUserRepository.EXPECT().InsertUser(u).Return(nil)
	err := user.InsertUser(mockUserRepository, u)
	if err != nil {
		t.Error("require no error")
	}
}
