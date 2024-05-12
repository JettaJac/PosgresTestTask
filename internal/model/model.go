package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// Command ... //TODO: возможно стелать как в осн , черезе репозиторий интерфейс В2-15.30
type Command struct {
	ID     int    `json:"id" db:"id"`
	Script string `json:"script" db:"script" validate:"required"` // возможно прикрутить валидацию
	Result string `json:"result" db:"result"`
	// TODO:   если успею прикрутить авторизацию /юзера и пароль
}

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
