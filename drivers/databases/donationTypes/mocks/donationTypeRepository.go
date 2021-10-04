package mocks
import (
	"context"
	donationtypes "finalproject-BE/business/donationTypes"

	"github.com/stretchr/testify/mock"
)

type DonationTypeRepository struct {
	mock.Mock
}

func (_m *DonationTypeRepository) AddDonationType(ctx context.Context, domain donationtypes.Domain) (donationtypes.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 donationtypes.Domain
	if rf, ok := ret.Get(0).(func(context.Context, donationtypes.Domain) donationtypes.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(donationtypes.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, donationtypes.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *DonationTypeRepository) GetDonationType(ctx context.Context) ([]donationtypes.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []donationtypes.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []donationtypes.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]donationtypes.Domain)
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

func (_m *DonationTypeRepository) GetDetailDonationType(ctx context.Context, id int) (donationtypes.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 donationtypes.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) donationtypes.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(donationtypes.Domain)
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

func (_m *DonationTypeRepository) UpdateDonationType(ctx context.Context, domain donationtypes.Domain) (donationtypes.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 donationtypes.Domain
	if rf, ok := ret.Get(0).(func(context.Context, donationtypes.Domain) donationtypes.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(donationtypes.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, donationtypes.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *DonationTypeRepository) DeleteDonationType(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}