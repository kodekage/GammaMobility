package dto

import (
	"time"
)

type CreateCustomerPaymentRequest struct {
	CustomerId           string    `json:"customer_id"`
	PaymentStatus        string    `json:"payment_status"`
	TransactionAmount    float64   `json:"transaction_amount"`
	TransactionDate      time.Time `json:"transaction_date"`
	TransactionReference string    `json:"transaction_reference"`
}
