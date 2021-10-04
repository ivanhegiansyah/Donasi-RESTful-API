package donations_test

import (
	"context"
	"errors"
	"finalproject-BE/business/donations"
	mockDonation "finalproject-BE/drivers/databases/donations/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	donationRepository mockDonation.DonationRepository
	donationService    donations.Usecase
	donationDomain     donations.Domain
)

func setup() {
	donationService = donations.NewDonationUsecase(&donationRepository, time.Hour*1)
	donationDomain = donations.Domain{
		Id:               1,
		UserId:           1,
		DonationTypeId:   1,
		DonationName:     "Bencana Gempa",
		Status:           "Active",
		ShortDescription: "Ini bencana gempa",
		GoalAmount:       1000000,
		CurrentAmount:    0,
		ExpiredDate:      "2022-15-05",
	}
}

func TestAddDonation(t *testing.T) {
	setup()
	donationRepository.On("AddDonation",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(donationDomain, nil).Once()
	t.Run("Test Case 1 | Valid Add Donation", func(t *testing.T) {
		_, err := donationService.AddDonation(context.Background(), donationDomain)

		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Invalid Add Donation Field Empty", func(t *testing.T) {
		_, err := donationService.AddDonation(context.Background(), donations.Domain{
			Id:               1,
			UserId:           1,
			DonationTypeId:   1,
			DonationName:     "Bencana Gempa",
			Status:           "Active",
			ShortDescription: "Ini bencana gempa",
			GoalAmount:       0,
			CurrentAmount:    0,
			ExpiredDate:      "",
		})
		assert.NotNil(t, err)
	})
}

func TestGetAllDonation(t *testing.T) {
	setup()
	donationRepository.On("GetAllDonation",
		mock.Anything).Return([]donations.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Get All Donation", func(t *testing.T) {
		_, err := donationService.GetAllDonation(context.Background())

		assert.Nil(t, err)
	})
}

func TestGetDetailDonation(t *testing.T) {
	setup()
	donationRepository.On("GetDetailDonation", mock.Anything, mock.AnythingOfType("int")).Return(donations.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Get Detail Donation", func(t *testing.T) {
		_, err := donationService.GetDetailDonation(context.Background(), 1)

		assert.Nil(t, err)
	})

	donationRepository.On("GetDetailDonation", mock.Anything, mock.AnythingOfType("int")).Return(donations.Domain{}, errors.New("Failed to Get Detail")).Once()
	t.Run("Test Case 2 | Invalid  Get Detail Donation", func(t *testing.T) {
		_, err := donationService.GetDetailDonation(context.Background(), 1)

		assert.Equal(t, err, errors.New("Failed to Get Detail"))
	})
}

func TestUpdateDonation(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Update Donation", func(t *testing.T) {
		donationRepository.On("UpdateDonation",
			mock.Anything,
			mock.AnythingOfType("Domain")).Return(donationDomain, nil).Once()

		_, err := donationService.UpdateDonation(context.Background(), donations.Domain{
			Id:               1,
			UserId:           1,
			DonationTypeId:   1,
			DonationName:     "Bencana Banjir",
			Status:           "Active",
			ShortDescription: "Ini bencana banjir",
			GoalAmount:       2000000,
			CurrentAmount:    0,
			ExpiredDate:      "2022-15-05",
		}, donationDomain.Id)

		assert.Nil(t, err)

		t.Run("Test Case 2 | Invalid Update Field Empty", func(t *testing.T) {
			_, err := donationService.UpdateDonation(context.Background(), donations.Domain{
				Id:               1,
				UserId:           1,
				DonationTypeId:   1,
				DonationName:     "Bencana Banjir",
				Status:           "Active",
				ShortDescription: "Ini bencana banjir",
				GoalAmount:       0,
				CurrentAmount:    0,
				ExpiredDate:      "",
			}, donationDomain.Id)
			assert.NotNil(t, err)
		})
	})
}

func TestDeleteDonation(t *testing.T) {
	setup()
	donationRepository.On("DeleteDonation", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
	t.Run("Test Case 1 | Valid Delete Donation", func(t *testing.T) {
		err := donationService.DeleteDonation(context.Background(), 1)

		assert.Nil(t, err)
	})

	donationRepository.On("DeleteDonation", mock.Anything, mock.AnythingOfType("int")).Return(errors.New("Failed to delete")).Once()
	t.Run("Test Case 2 | Invalid Delete Donation", func(t *testing.T) {
		err := donationService.DeleteDonation(context.Background(), 1)

		assert.Equal(t, err, errors.New("Failed to delete"))
	})
}
