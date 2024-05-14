package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// Command struct
type Command struct {
	ID     int    `json:"id" db:"id"`
	Script string `json:"script" db:"script" validate:"required"` // возможно прикрутить валидацию
	Result string `json:"result" db:"result"`
	// TODO:    авторизациz юзера и пароль
}

// ValidateJson validates json
func ValidateJson(req *Command) error {

	const op = "server.ValidateJson"

	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	_ = validate
	_ = err
	fmt.Println()
	return nil
}
