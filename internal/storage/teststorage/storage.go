package teststorage

import (
	// "main/internal/storage"
	"main/internal/model"
)

type Storage struct {
	storage *Storage
	commands map[string]*model.Command
}

func New() *Storage {
	return &Storage{}
}

// Create a new command 
func (r *Storage) SaveRunScript(c *model.Command) (int, error){
	// if err := u.Validate(); err != nil {
	// 	return err
	// }
	// if err := u.BeforeCreate(); err != nil {
	// 	return err
	// }

	r.commands[c.Script] = c
	c.ID = len(r.commands) /// Разобраться с инт64

	return c.ID, nil
}