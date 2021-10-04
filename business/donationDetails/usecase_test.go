package donationdetails_test

import (
	"context"
	"errors"
	donationdetails "finalproject-BE/business/donationDetails"
	mockDonationDetail "finalproject-BE/drivers/databases/donationDetails/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	donationDetailRepository mockDonationDetail.DonationDetailRepository
	donationDetailService    donationdetails.Usecase
	donationDetailDomain     donationdetails.Domain
)

func setup() {
	donationDetailService = donationdetails.NewDonationDetailUsecase(&donationDetailRepository, time.Hour*1)
	donationDetailDomain = donationdetails.Domain{
		Id:   1,
		DonationId: 1,
		Description: "Bencana banjir yang besar di suatu daerah",
	}
}

func TestAddDonationDetail(t *testing.T) {
	setup()
	donationDetailRepository.On("AddDonationDetail",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(donationDetailDomain, nil).Once()
	t.Run("Test Case 1 | Valid Add Donation Detail", func(t *testing.T) {
		_, err := donationDetailService.AddDonationDetail(context.Background(), donationDetailDomain, donationDetailDomain.DonationId)

		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Invalid Add Donation Detail Field Empty", func(t *testing.T) {
		_, err := donationDetailService.AddDonationDetail(context.Background(), donationdetails.Domain{
			Id:   1,
			DonationId: 1,
			Description: "",
		}, donationDetailDomain.DonationId)
		assert.NotNil(t, err)
	})
}

func TestUpdateDonationDetail(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Update Donation Detail", func(t *testing.T) {
		donationDetailRepository.On("UpdateDonationDetail",
			mock.Anything,
			mock.AnythingOfType("Domain")).Return(donationDetailDomain, nil).Once()

		_, err := donationDetailService.UpdateDonationDetail(context.Background(), donationdetails.Domain{
			Id:   1,
			DonationId: 1,
			Description: "Bencana Gempa di suatu daerah",
		}, donationDetailDomain.DonationId)

		assert.Nil(t, err)

		t.Run("Test Case 2 | Invalid Update Field Empty", func(t *testing.T) {
			_, err := donationDetailService.UpdateDonationDetail(context.Background(), donationdetails.Domain{
				Id:   1,
				DonationId: 1,
				Description: "",
			}, donationDetailDomain.DonationId)
			assert.NotNil(t, err)
		})
	})
}

func TestDeleteDonationDetail(t *testing.T) {
	setup()
	donationDetailRepository.On("DeleteDonationDetail", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
	t.Run("Test Case 1 | Valid Delete Donation Detail", func(t *testing.T) {
		err := donationDetailService.DeleteDonationDetail(context.Background(), 1)

		assert.Nil(t, err)
	})

	donationDetailRepository.On("DeleteDonationDetail", mock.Anything, mock.AnythingOfType("int")).Return(errors.New("Failed to delete")).Once()
	t.Run("Test Case 2 | Invalid Delete Donation Detail", func(t *testing.T) {
		err := donationDetailService.DeleteDonationDetail(context.Background(), 1)

		assert.Equal(t, err, errors.New("Failed to delete"))
	})
}
