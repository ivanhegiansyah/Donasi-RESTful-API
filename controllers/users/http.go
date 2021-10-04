package users

import (
	"finalproject-BE/business/users"
	"finalproject-BE/controllers"
	"finalproject-BE/controllers/users/requests"
	"finalproject-BE/controllers/users/responses"
	"fmt"
	"net/http"
	"strconv"

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
	user, error := userController.UserUseCase.Login(ctx, userLogin.ToDomainLogin())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainLogin(user))
}

func (userController UserController) Register(c echo.Context) error {
	fmt.Println("Register")
	userRegister := requests.UserRegister{}
	c.Bind(&userRegister)

	ctx := c.Request().Context()
	user, error := userController.UserUseCase.Register(ctx, userRegister.ToDomainRegister())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (userController UserController) GetAllUser(c echo.Context) error {
	fmt.Println("GetAll")
	ctx := c.Request().Context()
	user, error := userController.UserUseCase.GetAllUser(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromUserListDomain(user))
}

func (userController UserController) GetDetailUser(c echo.Context) error {
	fmt.Println("GetDetail")
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	user, error := userController.UserUseCase.GetDetailUser(ctx, id)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainDetail(user))
}

func (userController UserController) UpdateUser(c echo.Context) error {
	fmt.Println("Update")
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	userUpdate := requests.UserRegister{}
	c.Bind(&userUpdate)

	user, error := userController.UserUseCase.UpdateUser(ctx, userUpdate.ToDomainRegister(), id)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (userController UserController) DeleteUser(c echo.Context) error {
	fmt.Println("Delete")
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))


	err := userController.UserUseCase.DeleteUser(ctx, id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, nil)
}