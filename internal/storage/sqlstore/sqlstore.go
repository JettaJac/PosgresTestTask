package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct{
	db *sql.DB
}

fun New(storagePath string) (*Storage, error) { 
	const op "srorage.sqlstore.New"
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable") ///!!!перенести в другое место
	if err!= nil {
		retur nil, fmt.Errorf("%s: %s", op, err)
	}	

	return db, nil
}