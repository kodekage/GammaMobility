package entities

import "time"

type Transction struct {
	Id                    string    `db:"id"`
	Customer_id           string    `db:"customer_id"`
	Transaction_reference string    `db:"transaction_reference"`
	Amount                float32   `db:"amount"`
	Payment_status        string    `db:"payment_status"`
	CreatedAt             time.Time `db:"created_at"`
}
