package users

import (
	"finalproject-BE/business/users"
	"finalproject-BE/controllers"
	"finalproject-BE/controllers/users/requests"
	"finalproject-BE/controllers/users/responses"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(userUseCase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (userController UserController) Login(c echo.Context) error {
	fmt.Println("Login")
	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()
	user, error := userController.UserUseCase.Login(ctx, userLogin.Email, userLogin.Password)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (UserController UserController) Register(c echo.Context) error {
	fmt.Println("Register")
	userRegister := requests.UserRegister{}
	c.Bind(&userRegister)

	ctx := c.Request().Context()
	user, error := UserController.UserUseCase.Register(ctx, userRegister)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))

}