package repository

import "database/sql"

type CustomerRepository struct {
	DB *sql.DB
}
