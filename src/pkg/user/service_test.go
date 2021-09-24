package user_test

import (
	"testing"

	"github.com/bothonachiz/gotest/src/pkg/user"

	"github.com/bothonachiz/gotest/src/pkg/mocker"
	"github.com/bothonachiz/gotest/src/pkg/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	service          *user.Service
	repository       *mocks.UserRepository
	mockUserService  *mocks.UserService
	mockCallGetRepo  *mocker.MockCall
	mockCallGetLogin *mocker.MockCall
)

func callGetRepo() *mock.Call {
	return repository.On("GetRepo", mock.AnythingOfType("int"))
}

func callGetLogin() *mock.Call {
	return mockUserService.On("Login")
}

func beforeEach() {
	repository = &mocks.UserRepository{}
	service = user.NewService(repository)

	mockUserService = &mocks.UserService{}
	service.UserService = mockUserService

	mockCallGetRepo = mocker.NewMockCall(callGetRepo)
	mockCallGetRepo.NumberCallReturn(1, 0)

	mockCallGetLogin = mocker.NewMockCall(callGetLogin)
	mockCallGetLogin.NumberCallReturn(1, true)
}

func TestUser_GetInfo(t *testing.T) {
	beforeEachGetInfo := func() {
		beforeEach()
	}

	t.Run("should return false", func(t *testing.T) {
		beforeEachGetInfo()
		mockCallGetRepo.NumberCallReturn(1, 3).Once()
		mockCallGetRepo.NumberCallReturn(2, 4).Once()

		service.GetInfo()

		assert.True(t, true)
	})

	t.Run("should return false", func(t *testing.T) {
		beforeEachGetInfo()

		service.GetInfo()

		repository.AssertCalled(t, "GetRepo", 20)
	})

	t.Run("should login", func(t *testing.T) {
		beforeEachGetInfo()

		service.GetInfo()

		mockUserService.AssertCalled(t, "Login")
	})
}
