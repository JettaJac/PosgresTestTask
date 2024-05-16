package sqlstore

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
	"testing"
)

// TestDB creates a test database with the given name and runs the given
func TestDB(t *testing.T, databaseURL string) (*Storage, func(...string)) {
	t.Helper()
	op := "TestDB"

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		fmt.Printf("%s: %s", op, err)
		return nil, nil
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("%s: %s", op, err)
		return nil, nil
	}

	migrations(databaseURL)

	return &Storage{db}, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}
		db.Close()
	}
}
