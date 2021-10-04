package mocks

import (
	"context"
	"finalproject-BE/business/transactions"

	"github.com/stretchr/testify/mock"
)

type TransactionRepository struct {
	mock.Mock
}

func (_m *TransactionRepository) AddTransaction(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 transactions.Domain
	if rf, ok := ret.Get(0).(func(context.Context, transactions.Domain) transactions.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(transactions.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, transactions.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *TransactionRepository) GetAllTransaction(ctx context.Context) ([]transactions.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []transactions.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []transactions.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transactions.Domain)
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