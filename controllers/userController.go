package controllers

import (
	"finalproject-BE/config"
	"finalproject-BE/models/responses"
	"finalproject-BE/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

//CREATE user
func RegisterUserController(c echo.Context) error {
	userRegister := users.UserRegister{}
	c.Bind(&userRegister)

	//validation
	switch {
	case userRegister.Name == "":
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Nama harus diisi",
			Data:    nil,
		})
	case userRegister.Email == "":
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email harus diisi",
			Data:    nil,
		})
	case userRegister.Password == "":
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password harus diisi",
			Data:    nil,
		})
	}

	userDB := users.User{}
	userDB.Name = userRegister.Name
	userDB.Email = userRegister.Email
	userDB.Password = userRegister.Password
	userDB.Phone = userRegister.Phone
	userDB.Dob = userRegister.Dob

	result := config.DB.Create(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Terjadi kesalahan ketika input data user ke database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil registrasi",
		Data:    userDB,
	})
}

//READ user
func GetAllUserController(c echo.Context) error {
	users := []users.User{}

	result := config.DB.Find(&users)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data user dalam database",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data user",
		Data:    users,
	})
}

//READ 1 data user with query param
func GetOneUserController(c echo.Context) error {
	users := []users.User{}

	id, _ := strconv.Atoi(c.Param("id"))
	result := config.DB.First(&users, id)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data user dalam database",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data user",
		Data:    users,
	})
}

//UPDATE user
func UpdateUserController(c echo.Context) error {
	userUpdate := users.UserUpdate{}
	c.Bind(&userUpdate)

	users := []users.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	result := config.DB.First(&users, id)

	users[0].Name = userUpdate.Name
	users[0].Email = userUpdate.Email
	users[0].Password = userUpdate.Password
	users[0].Phone = userUpdate.Phone
	users[0].Dob = userUpdate.Dob
	config.DB.Save(&users)
	//jangan lupa tambahin jika pencarian tidak ketemu

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input memperbarui data user dalam database",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil memperbarui data user",
		Data:    users,
	})
}

//DELETE user
func DeleteUserController(c echo.Context) error {
	users := []users.User{}

	id, _ := strconv.Atoi(c.Param("id"))
	result := config.DB.Delete(&users, id)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input menghapus data userdalam database",
				Data:    nil,
			})
		}
	}
	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil menghapus data user",
		Data:    nil,
	})
}
