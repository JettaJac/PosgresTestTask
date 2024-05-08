package teststorage

import (
	// "main/internal/storage"
	"main/internal/model"
	// "net/http"
	// "fmt"
)

type Storage struct {
	// storage *Storage
	Commands map[string]*model.Command
	// Command map[string]int
}

func New() *Storage {
	return &Storage{
		// Command: make(map[string]int),
		Commands: make(map[string]*model.Command),
	}
}

// Create a new command
func (s *Storage) SaveRunScript(c *model.Command) (int, error) {

	s.Commands[c.Script] = c
	c.ID = len(s.Commands)

	return c.ID, nil
}

func (s *Storage) GetOneScript(req *model.Command) error {

	return nil
}

func (s *Storage) GetListCommands( /*req *model.Command, w http.ResponseWriter*/ ) ([]model.Command, error) {
	var commands []model.Command

	return commands, nil
}

func (s *Storage) DeleteCommand(id int) error {

	return nil
}
