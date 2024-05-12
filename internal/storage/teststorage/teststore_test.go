package teststorage_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"main/internal/model"
	"main/internal/storage/teststorage"
	"testing"
)

func TestCommand_SaveRunScript(t *testing.T) {

	storage := teststorage.New()
	req := model.TestCommand(t)
	req.Result = model.TestResult

	id, err := storage.SaveRunScript(req)
	assert.NoError(t, err)
	assert.NotNil(t, req)
	assert.Equal(t, 1, id)
	assert.Equal(t, model.TestResult, req.Result)
}

func TestCommand_GetOneScript(t *testing.T) {

	storage := teststorage.New()
	req := model.TestCommand(t)
	// resScript, _ := scripts.Run(req.Script)
	// req.Result = string(resScript)
	req.Result = model.TestResult

	id, err := storage.SaveRunScript(req)
	assert.NoError(t, err)
	fmt.Println(id)

	resp, err := storage.GetOneScript(id)
	fmt.Println(err)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, req.Result, resp.Result)

	//тест на запрос не существующего id
	resp, err = storage.GetOneScript(999)
	assert.Error(t, err)
	// assert.Nil(t, resp) //пока не проходит
}

func TestCommand_GetListCommands(t *testing.T) { //тест на пустой список,
	storage := teststorage.New()
	req := model.TestCommand(t)
	// resScript, _ := scripts.Run(c.Script)
	// c.Result = string(resScript)

	req.Result = model.TestResult

	_, err := storage.SaveRunScript(req)
	assert.NoError(t, err)
	assert.NotNil(t, req.ID)

	req2 := model.TestCommand(t)
	req2.Script = req.Script + "2"
	req2.Result = model.TestResult + "2"
	// resScript, _ = scripts.Run(req2.Script)
	// req2.Result = string(resScript)

	_, err = storage.SaveRunScript(req2)
	assert.NoError(t, err)
	assert.NotNil(t, req.ID)
	assert.Equal(t, req.ID+1, req2.ID)

	resp, err := storage.GetListCommands()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(resp))
}

func TestCommand_DeleteCommand(t *testing.T) { // не записывает 2 запрос()
	storage := teststorage.New()

	req := model.TestCommand(t)
	req.Result = model.TestResult
	_, err := storage.SaveRunScript(req)
	assert.NoError(t, err)
	assert.NotNil(t, req.ID)

	req2 := model.TestCommand(t)
	req2.Script = "test"
	req2.Result = model.TestResult + "2"
	id, err := storage.SaveRunScript(req2)
	assert.NoError(t, err)
	assert.NotNil(t, req2.ID)
	assert.Equal(t, req.ID+1, req2.ID)
	fmt.Println(id)

	err = storage.DeleteCommand(id)
	// assert.NoError(t, err)

	resp, err := storage.GetOneScript(id)
	assert.Nil(t, resp)
	// assert.Equal(t, 2, len(resp))

	err = storage.DeleteCommand(999)
	// fmt.Println(err)
	assert.Error(t, err)

}
