package mocks

import (
	"context"
	"finalproject-BE/business/donations"

	"github.com/stretchr/testify/mock"
)

type DonationRepository struct {
	mock.Mock
}

func (_m *DonationRepository) AddDonation(ctx context.Context, domain donations.Domain) (donations.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 donations.Domain
	if rf, ok := ret.Get(0).(func(context.Context, donations.Domain) donations.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(donations.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, donations.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *DonationRepository) GetAllDonation(ctx context.Context) ([]donations.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []donations.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []donations.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]donations.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *DonationRepository) GetDetailDonation(ctx context.Context, id int) (donations.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 donations.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) donations.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(donations.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *DonationRepository) UpdateDonation(ctx context.Context, domain donations.Domain) (donations.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 donations.Domain
	if rf, ok := ret.Get(0).(func(context.Context, donations.Domain) donations.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(donations.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, donations.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *DonationRepository) DeleteDonation(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}