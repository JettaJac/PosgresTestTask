package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	"main/internal/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func NewDBMY(storagePath string) (*Storage, error) { // !!! сторыдж патх заменить на инфу из конфига
	const op = "srorage.sqlstore.New"

	db, err := sql.Open("postgres", storagePath /*"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"*/) ///!!!перенести в другое место
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	if err := db.Ping(); err != nil {
		// fmt.Println("yyyyy")
		return nil, err
	}
	return &Storage{db}, nil
}

func (storage *Storage) CloseDB() {
	storage.db.Close()
}

func New(storagePath string) (*Storage, error) { // !!! сторыдж патх заменить на инфу из конфига
	const op = "srorage.sqlstore.New"
	db, err := sql.Open("postgres", storagePath /*"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"*/) ///!!!перенести в другое место
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	if err := db.Ping(); err != nil {
		// fmt.Println("yyyyy")
		return nil, err
	}
	// !!! база данных пока сама не создаеться, как и таблички

	// SELECT 'CREATE DATABASE mydb' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'mydb')\gexec
	/* возможно  id INTEGER PRIMARY KEY,*/
	// CREATE INDEX IF NOT EXISTS idx_name ON commands (name);
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS commands (
    id bigserial not null primary key, 
    name TEXT NOT NULL UNIQUE,
  	script TEXT NOT NULL);

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
	}

	return &Storage{db}, nil

}

func (s *Storage) SaveScript(urlTOSave, alias string) (int64, error) { //ште можно.нужно убрать
	const op = "storage.sqlstore.SaveScript"
	stmt, err := s.db.Prepare(`INSERT INTO commands (name,script) VALUES ($1,$2)`)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	res, err := stmt.Exec(urlTOSave, alias)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, storage.ErrURLExists) // !!!у  тузово по другому 47.49
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get laast insert id:  %w", op, err)
	}
	return id, nil
}

func (s *Storage) RunCommand(name, script string) error {
	// 	id := c.Param("id")
	//   row := db.QueryRow(`SELECT script FROM commands WHERE id = ?`, id)

	//   var script string
	//   if err := row.Scan(&script); err != nil {
	//    if err == sql.ErrNoRows {
	//     c.JSON(http.StatusNotFound, gin.H{"error": "Command not found"})
	//     return
	//    }
	//    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//    return
	//   }

	// 	result, err := exec.Command("bash", "-c", script).Output()
	// if err != nil {
	// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// return
	// }

	// command := Command{
	// ID: id,
	// Name: "command_" + id,
	// Script: script,
	// Result: string(result),
	// }

	// c.JSON(http.StatusOK, command)
	// })

	// r.Run(":8080"
	return nil
}

func (s *Storage) GetOneCommand(name string) (string, error) { //ште можно.нужно убрать jnlftn htpekmnfn
	const op = "storage.sqlstore.GetOneCommand"
	stmt, err := s.db.Prepare(`SELECT name FROM commands WHERE name = ?`)
	if err != nil {
		return "", fmt.Errorf("%s: prepare ststement: %w", op, err)
	}
	var result string
	err = stmt.QueryRow(name).Scan(&result)
	if errors.Is(err, sql.ErrNoRows) {
		return "", storage.ErrURLNotFound // !!!у  тузово по другому 47.49
	}

	if err != nil {
		return "", fmt.Errorf("%s: failed to get last insert id:  %w", op, err)
	}
	return result, nil
}

func (s *Storage) GetCommands(name string) (string, error) { //ште можно.нужно убрать jnlftn htpekmnfn
	const op = "storage.sqlstore.GetCommands"
	stmt, err := s.db.Prepare(`SELECT * FROM commands`)
	if err != nil {
		return "", fmt.Errorf("%s: prepare ststement: %w", op, err)
	}
	var result string
	err = stmt.QueryRow(name).Scan(&result)
	if errors.Is(err, sql.ErrNoRows) {
		return "", storage.ErrURLNotFound // !!!у  тузово по другому 47.49
	}

	if err != nil {
		return "", fmt.Errorf("%s: failed to get last insert id:  %w", op, err)
	}
	return result, nil
}

func (s *Storage) DeleteScript(name string) error { return nil }

// r.GET("/commands", func(c *gin.Context) {
//   rows, err := db.Query(`SELECT id, name, script FROM commands`)
//   if err != nil {
//    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//    return
//   }
//   defer rows.Close()

//   var commands []Command
//   for rows.Next() {
//    var command Command
//    if err := rows.Scan(&command.ID, &command.Name, &command.Script); err != nil {
//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//     return
//    }
//    commands = append(commands, command)
//   }

//   c.JSON(http.StatusOK, commands)
//  })
