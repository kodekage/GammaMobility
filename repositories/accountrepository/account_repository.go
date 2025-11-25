package accountrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodekage/gamma_mobility/entities"
	"github.com/kodekage/gamma_mobility/internal/errors"
	"github.com/kodekage/gamma_mobility/internal/logger"
)

type Repository interface {
	Save(a entities.Account) *errors.AppError
}

type accountrepository struct {
	sqlClient *pgxpool.Pool
}

var _ Repository = (*accountrepository)(nil)

func New(dbClient *pgxpool.Pool) Repository {
	return accountrepository{dbClient}
}

func (r accountrepository) Save(a entities.Account) *errors.AppError {
	query := `
		INSERT INTO accounts (customer_id, balance)
		VALUES ($1, $2)
		RETURNING id
	`

	var id string
	err := r.sqlClient.QueryRow(context.Background(), query, a.CustomerId, a.Balance).Scan(&id)
	if err != nil {
		logger.Error("Error while saving to account " + err.Error())
		return errors.NewUnexpectedError("Error " + err.Error())
	}

	return nil
}
