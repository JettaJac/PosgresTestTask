package sqlstore

// migrate create  -ext sql -dir migrations create_commands /*-seq init - команда файлы с миграциями
// migrate -path migrations -database "postgres://localhost/restapi_script?sslmode=disable" up

import (
	// "encoding/json"
	"database/sql"
	// "errors"
	"fmt"
	"main/internal/model"
	// "main/internal/storage"

	// "main/internal/storage"storage.sqlstore.sqlstore.New

	// "github.com/golang-migrate/migrate"
	_ "github.com/lib/pq"
	// "net/http"
)

var (
	Table = "commandsdb"
)

type Storage struct {
	db *sql.DB
}

func NewDB(storagePath string) (*Storage, error) {
	const op = "storage.sqlstore.sqlstore.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	/*// tmp
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS command2 (
	    id bigserial not null primary key,
	    name TEXT NOT NULL UNIQUE,
	  	script TEXT NOT NULL,
		result TEXT NOT NULL);

	`)
	if err != nil {
		fmt.Println("yyyyy", err)
		//   log.Fatal(err) // fmt.Errorf("%s: %s", op, err)
		// fmt.Errorf("%s: %s", op, err)
		fmt.Println("#{op}: #{err}")
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("#{op}: #{err}")
	}*/
	return &Storage{db}, nil
}

func (storage *Storage) CloseDB() {
	storage.db.Close()
}

func (s *Storage) SaveRunScript(req *model.Command) (int, error) { // CreateCommand( - название может такое  ..func GetCommands
	const op = "storage.sqlstore.SaveRunScript"
	query := fmt.Sprintf("INSERT INTO %s (script, result) VALUES ($1, $2) RETURNING id", Table)
	err := s.db.QueryRow(
		query,
		req.Script, req.Result,
	).Scan(&req.ID)
	if err != nil {
		return 0, fmt.Errorf("%s: %s", op, err) // лог не полный !!!
	}
	return req.ID, nil
}

func (s *Storage) GetOneScript(id int) (*model.Command, error) {
	const op = "storage.sqlstore.GetOneCommand"

	req := &model.Command{
		ID:     id,
		Script: "",
		Result: "",
	}

	query := fmt.Sprintf("SELECT script, result FROM %s WHERE id = $1", Table)
	err := s.db.QueryRow(query, req.ID).Scan(&req.Script, &req.Result)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err) // !!! посмотреть как здесь создаеться ошибка Command not found, которую мы потом проверяем
	}

	/*
		stmt, err := s.db.Prepare(`SELECT name FROM commands WHERE name = ?`)
		if err != nil {
			return "", fmt.Errorf("%s: prepare ststement: %w", op, err)
		}
		var result string
		err = stmt.QueryRow(id).Scan(&result)
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrURLNotFound // TODO: у  тузово по другому 47.49
		}

		if err != nil {
			return "", fmt.Errorf("%s: failed to get last insert id:  %w", op, err)
		}*/
	return req, nil
}

func (s *Storage) GetListCommands() ([]model.Command, error) {
	const op = "storage.sqlstore.GetListCommands"
	query := fmt.Sprintf("SELECT id, script, result  FROM %s", Table)
	// Выполнение SQL-запроса с db.Query
	var commands []model.Command
	// err := s.db.Select(&commands, query) !!! проверить сработает ли
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

	/*
		  /*
			rows, err := db.Query("SELECT id, name FROM users")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			// Итерация по результатам запроса
			for rows.Next() {
				var id int
				var name string
				if err := rows.Scan(&id, &name); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("ID: %d, Name: %s\n", id, name)
			}

			// Проверка ошибок после итерации
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}
		}
	*/

	return commands, nil
}

func (s *Storage) DeleteCommand(id int) error {
	const op = "storage.sqlstore.DeleteCommand"
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", Table)

	res, err := s.db.Exec(query, id)

	if err != nil { // !!! не выводит пока ошибку, если такой записи нет
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
