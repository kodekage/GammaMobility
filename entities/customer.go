package entities

import "time"

type Customer struct {
	Id        string `db:"id"`
	Name      string
	CreatedAt time.Time `db:"created_at"`
}
