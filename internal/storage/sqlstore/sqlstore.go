package sqlstore

// migrate create  -ext sql -dir migrations create_commands /*-seq init - команда файлы с миграциями
// migrate -path migrations -database "postgres://localhost/restapi_script?sslmode=disable" up

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"main/internal/model"
	//	github.com/golang-migrate/migrate v3.5.4+incompatible h1:R7OzwvCJTCgwapPCiX6DyBiu2czIUMDCB118gFTKTUA=
	//
	// github.com/golang-migrate/migrate v3.5.4+incompatible/go.mod h1:IsVUlFN5puWOmXrqjgGUfIRIbU7mr8oNBE2tyERd9Wk=
)

var (
	Table = "commandsdb"
)

type Storage struct {
	db *sql.DB
}

// NewDB create basa database for sql
func NewDB(storagePath string) (*Storage, error) {
	const op = "storage.sqlstore.sqlstore.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	migrations(storagePath, ".")

	return &Storage{db}, nil
}

// CloseDB close database
func (storage *Storage) CloseDB() {
	storage.db.Close()
}

// migrations create migrations
func migrations(host, path string) {
	qualy := fmt.Sprintf("file://%s/migrations", path)
	m, err := migrate.New(qualy, host)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()

	if errors.Is(err, errors.New("no change")) && err != nil {
		log.Fatal(err)
	}
}

// SaveCommand save command to database
func (s *Storage) SaveRunCommand(req *model.Command) (int, error) { // CreateCommand( - название может такое  ..func GetCommands
	const op = "storage.sqlstore.SaveRunCommand"
	query := fmt.Sprintf("INSERT INTO %s (script, result) VALUES ($1, $2) RETURNING id", Table)
	err := s.db.QueryRow(
		query,
		req.Script, req.Result,
	).Scan(&req.ID)

	if err != nil {
		return 0, fmt.Errorf("%s: %s", op, err)
	}

	return req.ID, nil
}

// GetOneCommand get command by id from database
func (s *Storage) GetOneCommand(id int) (*model.Command, error) {
	const op = "storage.sqlstore.GetOneCommand"

	req := &model.Command{
		ID:     id,
		Script: "",
		Result: "",
	}

	query := fmt.Sprintf("SELECT script, result FROM %s WHERE id = $1", Table)
	err := s.db.QueryRow(query, req.ID).Scan(&req.Script, &req.Result)

	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return req, nil
}

// GetListCommands get commands from database
func (s *Storage) GetListCommands() ([]model.Command, error) {
	const op = "storage.sqlstore.GetListCommands"
	query := fmt.Sprintf("SELECT id, script, result  FROM %s", Table)
	var commands []model.Command
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var command model.Command
		if err := rows.Scan(&command.ID, &command.Script, &command.Result); err != nil {
			return nil, fmt.Errorf("%s: %s", op, err)
		}
		commands = append(commands, command)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return commands, nil
}

// DeleteCommand delete command from database by ID
func (s *Storage) DeleteCommand(id int) error {
	const op = "storage.sqlstore.DeleteCommand"
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", Table)

	res, err := s.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: command with id %d not found", op, id)
	}

	return nil
}
