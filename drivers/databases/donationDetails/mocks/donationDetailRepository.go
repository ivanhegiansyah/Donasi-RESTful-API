package mocks
import (
	"context"
	donationdetails "finalproject-BE/business/donationDetails"

	"github.com/stretchr/testify/mock"
)

type DonationDetailRepository struct {
	mock.Mock
}

func (_m *DonationDetailRepository) AddDonationDetail(ctx context.Context, domain donationdetails.Domain, id int) (donationdetails.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 donationdetails.Domain
	if rf, ok := ret.Get(0).(func(context.Context, donationdetails.Domain) donationdetails.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(donationdetails.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, donationdetails.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *DonationDetailRepository) UpdateDonationDetail(ctx context.Context, domain donationdetails.Domain) (donationdetails.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 donationdetails.Domain
	if rf, ok := ret.Get(0).(func(context.Context, donationdetails.Domain) donationdetails.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(donationdetails.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, donationdetails.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *DonationDetailRepository) DeleteDonationDetail(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}