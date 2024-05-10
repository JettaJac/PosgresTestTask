package teststorage

import (
	// "main/internal/storage"
	"main/internal/model"
	"main/internal/storage"
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
func (s *Storage) SaveRunScript(req *model.Command) (int, error) {
	// req.Result = "Hello, World_Model_Test" !!!передается в тестах

	s.Commands[req.Script] = req
	req.ID = len(s.Commands)

	return req.ID, nil
}

func (s *Storage) GetOneScript(id int) (*model.Command, error) {
	count := 0
	req := &model.Command{
		ID:     0,
		Script: "",
		Result: "",
	}

	for _, v := range s.Commands {
		if v.ID == id {
			req = v
			count++
		}
	}
	if count < 0 {
		return nil, storage.ErrCommandNotFound
	}

	return req, nil
}

func (s *Storage) GetListCommands( /*req *model.Command, w http.ResponseWriter*/ ) ([]model.Command, error) {
	var commands []model.Command

	for _, v := range s.Commands {
		commands = append(commands, *v)
	}

	return commands, nil
}

func (s *Storage) DeleteCommand(id int) error {
	count := 0
	for _, v := range s.Commands {
		if v.ID == id {
			delete(s.Commands, v.Script)
			v = nil
			count++
		}
	}
	if count < 0 {
		return storage.ErrCommandNotFound
	}

	return nil
}
