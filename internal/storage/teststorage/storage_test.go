package teststorage_test

import (
	"main/internal/model"
	// "main/internal/storage"
	"main/internal/scripts"
	"main/internal/storage/teststorage"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestCommand_SaveRunScript(t *testing.T) {

	s := teststorage.New()
	c := model.TestCommand(t)
	resScript, _ := scripts.Run(c.Script) // возможно сразу готовое кинуть
	c.Result = string(resScript)
	// c.Result = "Hello, World_Test" // TODO:  зависит от Тест модели

	_, err := s.SaveRunScript(c)
	assert.NoError(t, err) //err
	assert.NotNil(t, c)

	// fmt.Println("storage ", s)
	// fmt.Println("test com ",c)
	// fmt.Println("резалт ", string(resScript), err)
}

func TestCommand_GetOneScript(t *testing.T) {
	s := teststorage.New()
	c := model.TestCommand(t)
	resScript, _ := scripts.Run(c.Script)
	c.Result = string(resScript)

	_, err := s.SaveRunScript(c)
	assert.NoError(t, err)
	/*id, */ err = s.GetOneScript(c) // возможно метод должен принимать только id,  а отдавать модель
	assert.NoError(t, err)
	assert.NotNil(t, c)
	// assert.Equal(t, 1, id)
}

func TestCommand_GetListCommands(t *testing.T) {
	s := teststorage.New()
	c := model.TestCommand(t)
	resScript, _ := scripts.Run(c.Script)
	c.Result = string(resScript)

	_, err := s.SaveRunScript(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	c2 := model.TestCommand(t)
	resScript, _ = scripts.Run(c2.Script)
	c2.Result = string(resScript)

	_, err = s.SaveRunScript(c2)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := s.GetListCommands()
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 2, len(res))

}
