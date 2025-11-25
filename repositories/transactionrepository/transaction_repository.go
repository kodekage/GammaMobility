package transactionrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodekage/gamma_mobility/entities"
	"github.com/kodekage/gamma_mobility/internal/errors"
	"github.com/kodekage/gamma_mobility/internal/logger"
)

type Repository interface {
	Save(t entities.Transction) *errors.AppError
}

type transactionrepository struct {
	sqlClient *pgxpool.Pool
}

var _ Repository = (*transactionrepository)(nil)

func New(dbClient *pgxpool.Pool) Repository {
	return transactionrepository{dbClient}
}

func (r transactionrepository) Save(e entities.Transction) *errors.AppError {
	defer r.sqlClient.Close()

	query := `
		INSERT INTO transactions (customer_id, transaction_reference, amount, payment_status)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id string
	err := r.sqlClient.QueryRow(context.Background(), query, e.CustomerId, e.TransactionReference, e.Amount, e.PaymentStatus).Scan(&id)
	if err != nil {
		logger.Error("Error while saving new transaction " + err.Error())
		return errors.NewUnexpectedError("Error " + err.Error())
	}

	return nil
}
