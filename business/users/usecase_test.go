package users_test

import (
	"context"
	"errors"
	_middleware "finalproject-BE/app/middlewares"
	"finalproject-BE/business/users"
	mockUser "finalproject-BE/drivers/databases/users/mocks"
	"finalproject-BE/helpers/encrypt"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepository mockUser.UserRepository
	userService    users.Usecase
	userDomain     users.Domain
)

func setup() {
	jwtConfig := &_middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}
	userService = users.NewUserUsecase(&userRepository, time.Hour*1, jwtConfig)
	userDomain = users.Domain{
		Id:       1,
		Name:     "ivan",
		Email:    "ivan@gmail.com",
		Password: "123",
		Phone:    "081238888354",
		Dob:      "2001-05-15",
		Token:    "123456",
	}
}

func TestLogin(t *testing.T) {
	setup()
	userDomain.Password, _ = encrypt.Hash(userDomain.Password)
	userRepository.On("Login",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(userDomain, nil).Once()

	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {

		user, err := userService.Login(context.Background(), users.Domain{
			Email:    "ivan@gmail.com",
			Password: "123",
		})

		assert.Nil(t, err)
		assert.Equal(t, "ivan", user.Name)
		assert.Equal(t, "ivan@gmail.com", user.Email)
	})
	t.Run("Test Case 2 | Invalid Email Empty", func(t *testing.T) {
		_, err := userService.Login(context.Background(), users.Domain{
			Email:    "",
			Password: "123",
		})

		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Invalid Password Empty", func(t *testing.T) {
		_, err := userService.Login(context.Background(), users.Domain{
			Email:    "ivan@gmail.com",
			Password: "",
		})

		assert.NotNil(t, err)
	})
}

func TestRegister(t *testing.T) {
	setup()
	userRepository.On("Register",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(userDomain, nil).Once()
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		_, err := userService.Register(context.Background(), userDomain)

		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Invalid Register Field Empty", func(t *testing.T) {
		_, err := userService.Register(context.Background(), users.Domain{
			Email:    "ivan@gmail.com",
			Password: "",
			Name:     "ivan",
			Phone:    "0891828282",
			Dob:      "",
		})
		assert.NotNil(t, err)
	})
}

func TestGetAllUser(t *testing.T) {
	setup()
	userRepository.On("GetAllUser",
	mock.Anything).Return([]users.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Get All User", func(t *testing.T) {
		_, err := userService.GetAllUser(context.Background())
		assert.Nil(t, err)
	})

}

func TestGetDetailUser(t *testing.T) {
	setup()
	userRepository.On("GetDetailUser", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Get Detail User", func(t *testing.T) {
		_, err := userService.GetDetailUser(context.Background(), 1)

		assert.Nil(t, err)
	})

	userRepository.On("GetDetailUser", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("Failed to Get Detail")).Once()
	t.Run("Test Case 2 | Invalid  Get Detail User", func(t *testing.T) {
		_, err := userService.GetDetailUser(context.Background(), 1)

		assert.Equal(t, err, errors.New("Failed to Get Detail"))
	})
}

func TestUpdateUser(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Update User", func(t *testing.T) {
		userRepository.On("UpdateUser",
			mock.Anything,
			mock.AnythingOfType("Domain")).Return(userDomain, nil).Once()

		_, err := userService.UpdateUser(context.Background(), users.Domain{
			Name:     "ivanhegiansyah",
			Email:    "ivanhegiansyah@gmail.com",
			Password: "123",
			Phone:    "081238888354",
			Dob:      "2001-05-15",
		}, userDomain.Id)

		assert.Nil(t, err)

		t.Run("Test Case 2 | Invalid Update Field Empty", func(t *testing.T) {
			_, err := userService.UpdateUser(context.Background(), users.Domain{
				Email:    "ivan@gmail.com",
				Password: "",
				Name:     "ivan",
				Phone:    "0891828282",
				Dob:      "",
			}, userDomain.Id)
			assert.NotNil(t, err)
		})
	})
}

func TestDeleteUser(t *testing.T) {
	setup()
	userRepository.On("DeleteUser", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
	t.Run("Test Case 1 | Valid Delete User", func(t *testing.T) {
		err := userService.DeleteUser(context.Background(), 1)

		assert.Nil(t, err)
	})

	userRepository.On("DeleteUser", mock.Anything, mock.AnythingOfType("int")).Return(errors.New("Failed to delete")).Once()
	t.Run("Test Case 2 | Invalid Delete User", func(t *testing.T) {
		err := userService.DeleteUser(context.Background(), 1)

		assert.Equal(t, err, errors.New("Failed to delete"))
	})
}
