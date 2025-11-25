package transactionrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodekage/gamma_mobility/entities"
	"github.com/kodekage/gamma_mobility/internal/logger"
)

type Repository interface {
	Save(t entities.Transction) error
}

type transactionrepository struct {
	sqlClient *pgxpool.Pool
}

var _ Repository = (*transactionrepository)(nil)

func New(dbClient *pgxpool.Pool) Repository {
	return transactionrepository{dbClient}
}

func (r transactionrepository) Save(e entities.Transction) error {
	query := `
		INSERT INTO transactions (customer_id, transaction_reference, amount, payment_status, transaction_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var id string
	err := r.sqlClient.QueryRow(context.Background(), query, e.CustomerId, e.TransactionReference, e.Amount, e.PaymentStatus, e.TransactionDate).Scan(&id)
	if err != nil {
		logger.Error("Error while saving new transaction " + err.Error())
		return err
	}

	return nil
}
