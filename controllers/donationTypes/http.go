package donationtypes

import (
	donationtypes "finalproject-BE/business/donationTypes"
	"finalproject-BE/controllers"
	"finalproject-BE/controllers/donationTypes/requests"
	"finalproject-BE/controllers/donationTypes/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DonationTypeController struct {
	DonationTypeUseCase donationtypes.Usecase
}

func NewDonationTypeController(donationTypeUseCase donationtypes.Usecase) *DonationTypeController {
	return &DonationTypeController{
		DonationTypeUseCase: donationTypeUseCase,
	}
}

func (donationTypeController DonationTypeController) AddDonationType(c echo.Context) error {
	fmt.Println("Add")
	donationTypeAdd := requests.DonationTypeAdd{}
	c.Bind(&donationTypeAdd)

	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	donationType, error := donationTypeController.DonationTypeUseCase.AddDonationType(ctx, donationTypeAdd.ToDomainAdd(), id)
	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(donationType))

}