package paymentservice

import (
	"github.com/kodekage/gamma_mobility/dto"
	"github.com/kodekage/gamma_mobility/repositories/customerrepository"
)

type Service interface {
	ProcessPayment(data dto.CreateCustomerPaymentRequest) error
}

type paymentService struct {
	repo customerrepository.Repository
}

var _ Service = (*paymentService)(nil)

func New(repo customerrepository.Repository) Service {
	return paymentService{repo}
}

func (p paymentService) ProcessPayment(data dto.CreateCustomerPaymentRequest) error {
	return nil
}
