package sqlstore

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct{
	db *sql.DB
}

func New(storagePath string) (*Storage, error) { 
	const op = "srorage.sqlstore.New"
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable") ///!!!перенести в другое место
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}	

	return &Storage{db}, nil
}