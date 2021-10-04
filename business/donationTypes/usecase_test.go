package donationtypes_test

import (
	"context"
	"errors"
	donationtypes "finalproject-BE/business/donationTypes"
	mockDonationType "finalproject-BE/drivers/databases/donationTypes/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	donationTypeRepository mockDonationType.DonationTypeRepository
	donationTypeService    donationtypes.Usecase
	donationTypeDomain     donationtypes.Domain
)

func setup() {
	donationTypeService = donationtypes.NewDonationTypeUsecase(&donationTypeRepository, time.Hour*1)
	donationTypeDomain = donationtypes.Domain{
		Id:   1,
		Name: "Bencana",
	}
}

func TestAddDonationType(t *testing.T) {
	setup()
	donationTypeRepository.On("AddDonationType",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(donationTypeDomain, nil).Once()
	t.Run("Test Case 1 | Valid Add Donation Type", func(t *testing.T) {
		_, err := donationTypeService.AddDonationType(context.Background(), donationTypeDomain)

		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Invalid Add Donation Type Field Empty", func(t *testing.T) {
		_, err := donationTypeService.AddDonationType(context.Background(), donationtypes.Domain{
			Id:   1,
			Name: "",
		})
		assert.NotNil(t, err)
	})
}

func TestGetDonationType(t *testing.T) {
	setup()
	donationTypeRepository.On("GetDonationType",
		mock.Anything).Return([]donationtypes.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Get Donation Types", func(t *testing.T) {
		_, err := donationTypeService.GetDonationType(context.Background())

		assert.Nil(t, err)
	})
}

func TestGetDetailDonationTypeType(t *testing.T) {
	setup()
	donationTypeRepository.On("GetDetailDonationType", mock.Anything, mock.AnythingOfType("int")).Return(donationtypes.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Get Detail Donation Type", func(t *testing.T) {
		_, err := donationTypeService.GetDetailDonationType(context.Background(), 1)

		assert.Nil(t, err)
	})

	donationTypeRepository.On("GetDetailDonationType", mock.Anything, mock.AnythingOfType("int")).Return(donationtypes.Domain{}, errors.New("Failed to Get Detail")).Once()
	t.Run("Test Case 2 | Invalid  Get Detail Donation Type", func(t *testing.T) {
		_, err := donationTypeService.GetDetailDonationType(context.Background(), 1)

		assert.Equal(t, err, errors.New("Failed to Get Detail"))
	})
}

func TestUpdateDonationType(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Update Donation Type", func(t *testing.T) {
		donationTypeRepository.On("UpdateDonationType",
			mock.Anything,
			mock.AnythingOfType("Domain")).Return(donationTypeDomain, nil).Once()

		_, err := donationTypeService.UpdateDonationType(context.Background(), donationtypes.Domain{
			Id:   1,
			Name: "Pendidikan",
		}, donationTypeDomain.Id)

		assert.Nil(t, err)

		t.Run("Test Case 2 | Invalid Update Field Empty", func(t *testing.T) {
			_, err := donationTypeService.UpdateDonationType(context.Background(), donationtypes.Domain{
				Id:   1,
				Name: "",
			}, donationTypeDomain.Id)
			assert.NotNil(t, err)
		})
	})
}

func TestDeleteDonationType(t *testing.T) {
	setup()
	donationTypeRepository.On("DeleteDonationType", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()
	t.Run("Test Case 1 | Valid Delete Donation Type", func(t *testing.T) {
		err := donationTypeService.DeleteDonationType(context.Background(), 1)

		assert.Nil(t, err)
	})

	donationTypeRepository.On("DeleteDonationType", mock.Anything, mock.AnythingOfType("int")).Return(errors.New("Failed to delete")).Once()
	t.Run("Test Case 2 | Invalid Delete Donation Type", func(t *testing.T) {
		err := donationTypeService.DeleteDonationType(context.Background(), 1)

		assert.Equal(t, err, errors.New("Failed to delete"))
	})
}
