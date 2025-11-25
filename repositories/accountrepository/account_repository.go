package accountrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodekage/gamma_mobility/entities"
	"github.com/kodekage/gamma_mobility/internal/logger"
)

type Repository interface {
	Save(a entities.Account) error
	GetByCustomerId(id string) (*entities.Account, error)
}

type accountrepository struct {
	sqlClient *pgxpool.Pool
}

var _ Repository = (*accountrepository)(nil)

func New(dbClient *pgxpool.Pool) Repository {
	return accountrepository{dbClient}
}

func (r accountrepository) Save(a entities.Account) error {
	query := `
		INSERT INTO accounts (customer_id, balance)
		VALUES ($1, $2)
		RETURNING id
	`

	var id string
	err := r.sqlClient.QueryRow(context.Background(), query, a.CustomerId, a.Balance).Scan(&id)
	if err != nil {
		logger.Error("Error while saving to account " + err.Error())
		return err
	}

	return nil
}

func (r accountrepository) GetByCustomerId(id string) (*entities.Account, error) {
	var account entities.Account

	query := `SELECT id, customer_id, balance, created_at FROM accounts WHERE customer_id=$1`
	err := r.sqlClient.QueryRow(context.Background(), query, id).Scan(
		&account.Id,
		&account.CustomerId,
		&account.Balance,
		&account.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
