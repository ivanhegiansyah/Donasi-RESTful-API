package requests

import paymentmethods "finalproject-BE/business/paymentMethods"

type PaymentMethodAdd struct {
	MethodName string `json:"methodName"`
}

func (paymentMethod *PaymentMethodAdd) ToDomainAdd() paymentmethods.Domain {
	return paymentmethods.Domain{
		MethodName: paymentMethod.MethodName,
	}
}