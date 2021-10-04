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
	donationType, error := donationTypeController.DonationTypeUseCase.AddDonationType(ctx, donationTypeAdd.ToDomainAdd())
	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(donationType))

}

func (donationTypeController DonationTypeController) GetDonationType(c echo.Context) error {
	fmt.Println("GetAll")
	ctx := c.Request().Context()
	donationTypeAdd, error := donationTypeController.DonationTypeUseCase.GetDonationType(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomain(donationTypeAdd))
}

func (donationTypeController DonationTypeController) GetDetailDonationType(c echo.Context) error {
	fmt.Println("GetDetail")
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	donationTypeAdd, error := donationTypeController.DonationTypeUseCase.GetDetailDonationType(ctx, id)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainDetail(donationTypeAdd))
}

func (donationTypeController DonationTypeController) UpdateDonationType(c echo.Context) error {
	fmt.Println("Update")
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	donationTypeAdd := requests.DonationTypeAdd{}
	c.Bind(&donationTypeAdd)

	donation, error := donationTypeController.DonationTypeUseCase.UpdateDonationType(ctx, donationTypeAdd.ToDomainAdd(), id)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(donation))
}

func (donationTypeController DonationTypeController) DeleteDonationType(c echo.Context) error {
	fmt.Println("Delete")
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))


	err := donationTypeController.DonationTypeUseCase.DeleteDonationType(ctx, id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, nil)
}