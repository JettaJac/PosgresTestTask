package teststorage_test

import (
	"main/internal/model"
	// "main/internal/storage"
	"main/internal/scripts"
	"main/internal/storage/teststorage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_SaveRunScript(t *testing.T) {

	s := teststorage.New()
	c := model.TestCommand(t)
	resScript, _ := scripts.Run(c.Script) // возможно сразу готовое кинуть
	c.Result = string(resScript)
	c.Result = "Hello, World_Test" // TODO:  зависит от Тест модели

	_, err := s.SaveRunScript(c)
	assert.NoError(t, err) //err
	assert.NotNil(t, c)
}