package donationdetails

import (
	donationdetails "finalproject-BE/business/donationDetails"
	"finalproject-BE/controllers"
	"finalproject-BE/controllers/donationDetails/requests"
	"finalproject-BE/controllers/donationDetails/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DonationDetailController struct {
	DonationDetailUseCase donationdetails.Usecase
}

func NewDonationDetailController(donationDetailUseCase donationdetails.Usecase) *DonationDetailController {
	return &DonationDetailController{
		DonationDetailUseCase: donationDetailUseCase,
	}
}

func (donationDetailController DonationDetailController) AddDonationDetail(c echo.Context) error {
	fmt.Println("Add")
	donationDetailAdd := requests.DonationDetailAdd{}
	c.Bind(&donationDetailAdd)

	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	donationDetail, error := donationDetailController.DonationDetailUseCase.AddDonationDetail(ctx, donationDetailAdd.ToDomainAdd(), id)
	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(donationDetail))

}

func (donationDetailController DonationDetailController) UpdatedonationDetail(c echo.Context) error {
	fmt.Println("Update")
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	donationDetailUpdate := requests.DonationDetailAdd{}
	c.Bind(&donationDetailUpdate)

	donation, error := donationDetailController.DonationDetailUseCase.UpdateDonationDetail(ctx, donationDetailUpdate.ToDomainAdd(), id)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(donation))
}

func (donationDetailController DonationDetailController) DeleteDonationDetail(c echo.Context) error {
	fmt.Println("Delete")
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))


	err := donationDetailController.DonationDetailUseCase.DeleteDonationDetail(ctx, id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, nil)
}