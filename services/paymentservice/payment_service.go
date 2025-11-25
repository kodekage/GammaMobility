package paymentservice

import (
	"github.com/kodekage/gamma_mobility/dto"
	"github.com/kodekage/gamma_mobility/entities"
	"github.com/kodekage/gamma_mobility/internal/logger"
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
	_, err := p.customerRepository.GetCustomerByID(data.CustomerId)
	if err != nil {
		logger.Error("Customer Not Found " + err.Error())
		return err
	}

	// NOTE: Ideally every transaction reference should be validated against the originiating payment provider
	// but its omitted assuming transactions are verified.

	// Record Transaction
	transaction := entities.Transction{
		CustomerId:           data.CustomerId,
		TransactionReference: data.TransactionReference,
		Amount:               float32(data.TransactionAmount),
		PaymentStatus:        data.PaymentStatus,
		TransactionDate:      data.TransactionDate,
	}
	if err := p.transactionrepository.Save(transaction); err != nil {
		logger.Error(err.Error())
		return err
	}

	// Update customer account
	account, err := p.accountrepository.GetByCustomerId(data.CustomerId)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	new_account := entities.Account{
		CustomerId:      data.CustomerId,
		Balance:         account.Balance + float32(data.TransactionAmount),
		TotalAssetValue: account.TotalAssetValue,
		CreatedAt:       account.CreatedAt,
	}
	if err := p.accountrepository.Save(new_account); err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
