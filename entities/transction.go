package entities

import "time"

type Transction struct {
	Id                   string    `db:"id"`
	CustomerId           string    `db:"customer_id"`
	TransactionReference string    `db:"transaction_reference"`
	Amount               float32   `db:"amount"`
	PaymentStatus        string    `db:"payment_status"`
	TransactionDate      time.Time `db:"transaction_date"`
}
