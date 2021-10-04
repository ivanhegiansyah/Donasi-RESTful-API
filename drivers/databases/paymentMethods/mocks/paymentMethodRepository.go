package mocks

import (
	"context"
	paymentmethods "finalproject-BE/business/paymentMethods"

	"github.com/stretchr/testify/mock"
)

type PaymentMethodRepository struct {
	mock.Mock
}

func (_m *PaymentMethodRepository) AddPaymentMethod(ctx context.Context, domain paymentmethods.Domain) (paymentmethods.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 paymentmethods.Domain
	if rf, ok := ret.Get(0).(func(context.Context, paymentmethods.Domain) paymentmethods.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(paymentmethods.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, paymentmethods.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *PaymentMethodRepository) GetAllPaymentMethod(ctx context.Context) ([]paymentmethods.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []paymentmethods.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []paymentmethods.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]paymentmethods.Domain)
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