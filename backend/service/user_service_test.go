package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/dimgsg9/booker_proto/backend/model"
	"github.com/dimgsg9/booker_proto/backend/model/apperrors"
	"github.com/dimgsg9/booker_proto/backend/model/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := &model.User{
			UID:   uid,
			Email: "bob@bob.com",
			Name:  "Bobby Bobson",
		}

		mockUserRepository := new(mocks.MockUserRepository)
		us := NewUserService(&USConfig{
			UserRepository: mockUserRepository,
		})
		mockUserRepository.On("FindByID", mock.Anything, uid).Return(mockUserResp, nil)

		ctx := context.TODO()
		u, err := us.Get(ctx, uid)

		assert.NoError(t, err)
		assert.Equal(t, u, mockUserResp)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserRepository := new(mocks.MockUserRepository)
		us := NewUserService(&USConfig{
			UserRepository: mockUserRepository,
		})

		mockUserRepository.On("FindByID", mock.Anything, uid).Return(nil, fmt.Errorf("Some error down the call chain"))

		ctx := context.TODO()
		u, err := us.Get(ctx, uid)

		assert.Nil(t, u)
		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestSignup(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUser := &model.User{
			Email:    "bob@bob.com",
			Password: "howdyhoneighbor!",
		}

		mockUserRepository := new(mocks.MockUserRepository)
		us := NewUserService(&USConfig{
			UserRepository: mockUserRepository,
		})

		// We can use Run method to modify the user when the Create method is called.
		//  We can then chain on a Return method to return no error
		mockUserRepository.
			On("Create", mock.AnythingOfType("*context.emptyCtx"), mockUser).
			Run(func(args mock.Arguments) {
				userArg := args.Get(1).(*model.User) // arg 0 is context, arg 1 is *User
				userArg.UID = uid
			}).Return(nil)

		ctx := context.TODO()
		err := us.Signup(ctx, mockUser)

		assert.NoError(t, err)

		// assert user now has a userID
		assert.Equal(t, uid, mockUser.UID)

		mockUserRepository.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockUser := &model.User{
			Email:    "bob@bob.com",
			Password: "howdyhoneighbor!",
		}

		mockUserRepository := new(mocks.MockUserRepository)
		us := NewUserService(&USConfig{
			UserRepository: mockUserRepository,
		})

		mockErr := apperrors.NewConflict("email", mockUser.Email)

		// We can use Run method to modify the user when the Create method is called.
		//  We can then chain on a Return method to return no error
		mockUserRepository.
			On("Create", mock.AnythingOfType("*context.emptyCtx"), mockUser).
			Return(mockErr)

		ctx := context.TODO()
		err := us.Signup(ctx, mockUser)

		// assert error is error we response with in mock
		assert.EqualError(t, err, mockErr.Error())

		mockUserRepository.AssertExpectations(t)
	})
}
