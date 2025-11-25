package customerrepository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodekage/gamma_mobility/entities"
)

type Repository interface {
	GetCustomerByID(id string) (*entities.Customer, error)
}

type customerRepository struct {
	sqlClient *pgxpool.Pool
}

var _ Repository = (*customerRepository)(nil)

func New(dbClient *pgxpool.Pool) Repository {
	return customerRepository{dbClient}
}

func (r customerRepository) GetCustomerByID(id string) (*entities.Customer, error) {
	var customer entities.Customer
	query := `SELECT id, name, created_at FROM customers WHERE id=$1`
	err := r.sqlClient.QueryRow(context.Background(), query, id).Scan(
		&customer.Id,
		&customer.Name,
		&customer.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
