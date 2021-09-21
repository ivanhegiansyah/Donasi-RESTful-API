package controllers

import (
	"finalproject-BE/config"
	"finalproject-BE/models/donations"
	"finalproject-BE/models/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

//CREATE donation
func AddDonationrController(c echo.Context) error {
	donationRegister := donations.DonationRegister{}
	c.Bind(&donationRegister)

	userId, _ := strconv.Atoi(c.Param("userid"))

	//validation
	switch {
	case donationRegister.Donation_name == "":
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Nama harus diisi",
			Data:    nil,
		})
	case donationRegister.Short_description == "":
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "deskripsi singkat harus diisi",
			Data:    nil,
		})
	case donationRegister.Goal_amount == 0:
		return c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Target donasi harus diisi",
			Data:    nil,
		})
	default:
		{
			donationRegister.Status = "Active"
			donationRegister.Current_amount = 0
		}
	}

	//parsing string to time
	var date, _ = time.Parse(time.RFC822, donationRegister.Expired_date)

	donationDB := donations.Donation{}
	donationDB.Donation_name = donationRegister.Donation_name
	donationDB.Status = donationRegister.Status
	donationDB.Short_description = donationRegister.Short_description
	donationDB.Goal_amount = donationRegister.Goal_amount
	donationDB.Current_amount = donationRegister.Current_amount
	donationDB.Expired_date = date
	donationDB.UserID = uint(userId)

	result := config.DB.Create(&donationDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Terjadi kesalahan ketika input data penggalangan dana ke database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil membuat penggalangan dana",
		Data:    donationDB,
	})
}

//READ user
func GetAllDonationController(c echo.Context) error {
	donations := []donations.Donation{}

	userId, _ := strconv.Atoi(c.Param("userid"))
	result := config.DB.Where("user_id = ?", userId).Find(&donations)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data penggalangan dana dalam database",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data penggalangan dana",
		Data:    donations,
	})
}

//READ 1 data user with query param
func GetOneDonationController(c echo.Context) error {
	donations := []donations.Donation{}

	userId, _ := strconv.Atoi(c.Param("userid"))
	donationId, _ := strconv.Atoi(c.Param("donationid"))
	result := config.DB.Where("user_id = ? AND id = ?", userId, donationId).First(&donations)

	//blm validasi id user

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data penggalangan dana dalam database",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan penggalangan dana",
		Data:    donations,
	})
}

//UPDATE user
func UpdateDonationController(c echo.Context) error {
	donationUpdate := donations.DonationUpdate{}
	c.Bind(&donationUpdate)

	donations := []donations.Donation{}
	userId, _ := strconv.Atoi(c.Param("userid"))
	donationId, _ := strconv.Atoi(c.Param("donationid"))
	result := config.DB.Where("user_id = ? AND id = ?", userId, donationId).First(&donations)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input memperbarui data penggalangan dana dalam database",
				Data:    nil,
			})
		}
	}

	//parsing string to time
	var date, _ = time.Parse(time.RFC822, donationUpdate.Expired_date)

	donations[0].Donation_name = donationUpdate.Donation_name
	donations[0].Short_description = donationUpdate.Short_description
	donations[0].Goal_amount = donationUpdate.Goal_amount
	donations[0].Expired_date = date
	config.DB.Save(&donations)
	//jangan lupa tambahin jika pencarian tidak ketemu

	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil memperbarui data penggalangan dana",
		Data:    donations,
	})
}

//DELETE user
func DeleteDonationController(c echo.Context) error {
	donations := []donations.Donation{}

	userId, _ := strconv.Atoi(c.Param("userid"))
	donationId, _ := strconv.Atoi(c.Param("donationid"))
	result := config.DB.Where("user_id = ? AND id = ?", userId, donationId).Delete(&donations)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, responses.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input menghapus data penggalangan dana dalam database",
				Data:    nil,
			})
		}
	}
	return c.JSON(http.StatusOK, responses.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil menghapus data penggalangan dana",
		Data:    nil,
	})
}
