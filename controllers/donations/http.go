package donations

import (
	"finalproject-BE/business/donations"
	"finalproject-BE/controllers"
	"finalproject-BE/controllers/donations/requests"
	"finalproject-BE/controllers/donations/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DonationController struct {
	DonationUseCase donations.Usecase
}

func NewDonationController(donationUseCase donations.Usecase) *DonationController {
	return &DonationController{
		DonationUseCase: donationUseCase,
	}
}

func (donationController DonationController) AddDonation(c echo.Context) error {
	fmt.Println("Add")
	donationAdd := requests.DonationAdd{}
	c.Bind(&donationAdd)
	ctx := c.Request().Context()
	donation, error := donationController.DonationUseCase.AddDonation(ctx, donationAdd.ToDomainAdd())
	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(donation))

}

func (donationController DonationController) GetAllDonation(c echo.Context) error {
	fmt.Println("GetAll")
	ctx := c.Request().Context()
	donation, error := donationController.DonationUseCase.GetAllDonation(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDonationListDomain(donation))
}

func (donationController DonationController) GetDetailDonation(c echo.Context) error {
	fmt.Println("GetDetail")
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	donation, error := donationController.DonationUseCase.GetDetailDonation(ctx, id)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainDetail(donation))
}

func (donationController DonationController) Updatedonation(c echo.Context) error {
	fmt.Println("Update")
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	donationUpdate := requests.DonationAdd{}
	c.Bind(&donationUpdate)

	donation, error := donationController.DonationUseCase.UpdateDonation(ctx, donationUpdate.ToDomainAdd(), id)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(donation))
}

func (donationController DonationController) DeleteDonation(c echo.Context) error {
	fmt.Println("Delete")
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))


	err := donationController.DonationUseCase.DeleteDonation(ctx, id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, nil)
}

