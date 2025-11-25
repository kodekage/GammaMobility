package paymentservice

import (
	"github.com/kodekage/gamma_mobility/dto"
	"github.com/kodekage/gamma_mobility/repositories/accountrepository"
	"github.com/kodekage/gamma_mobility/repositories/customerrepository"
	"github.com/kodekage/gamma_mobility/repositories/transactionrepository"
	"github.com/kodekage/gamma_mobility/utils"
)

var (
	sqlClient = utils.SqlClient()
)

type Service interface {
	ProcessPayment(data dto.CreateCustomerPaymentRequest) error
}

type paymentService struct {
	customerRepository    customerrepository.Repository
	transactionrepository transactionrepository.Repository
	accountrepository     accountrepository.Repository
}

var _ Service = (*paymentService)(nil)

func New() Service {
	return paymentService{
		customerRepository:    customerrepository.New(sqlClient),
		transactionrepository: transactionrepository.New(sqlClient),
		accountrepository:     accountrepository.New(sqlClient)}
}

func (p paymentService) ProcessPayment(data dto.CreateCustomerPaymentRequest) error {
	return nil
}
