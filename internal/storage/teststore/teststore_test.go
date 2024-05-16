package teststore_test

import (
	"main/internal/model"
	"main/internal/storage/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests for map

// TestCommand_SaveRunCommand ...
func TestCommand_SaveRunCommand(t *testing.T) {

	storage := teststore.New()
	req := model.TestCommand(t)
	req.Result = model.TestResult

	id, err := storage.SaveRunCommand(req)
	assert.NoError(t, err)
	assert.NotNil(t, req)
	assert.Equal(t, 1, id)
	assert.Equal(t, model.TestResult, req.Result)
}

// TestCommand_GetOneCommand...
func TestCommand_GetOneCommand(t *testing.T) {

	storage := teststore.New()
	req := model.TestCommand(t)
	req.Result = model.TestResult

	id, err := storage.SaveRunCommand(req)
	assert.NoError(t, err)

	resp, err := storage.GetOneCommand(id)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, req.Result, resp.Result)

	resp, err = storage.GetOneCommand(999)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

// TestCommand_GetListCommands...
func TestCommand_GetListCommands(t *testing.T) { //тест на пустой список,
	storage := teststore.New()
	req := model.TestCommand(t)

	req.Result = model.TestResult

	_, err := storage.SaveRunCommand(req)
	assert.NoError(t, err)
	assert.NotNil(t, req.ID)

	req2 := model.TestCommand(t)
	req2.Script = req.Script + "2"
	req2.Result = model.TestResult + "2"

	_, err = storage.SaveRunCommand(req2)
	assert.NoError(t, err)
	assert.NotNil(t, req.ID)
	assert.Equal(t, req.ID+1, req2.ID)

	resp, err := storage.GetListCommands()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(resp))
}

// / TestCommand_DeleteCommand...
func TestCommand_DeleteCommand(t *testing.T) { // не записывает 2 запрос()
	storage := teststore.New()

	req := model.TestCommand(t)
	req.Result = model.TestResult
	_, err := storage.SaveRunCommand(req)
	assert.NoError(t, err)
	assert.NotNil(t, req.ID)

	req2 := model.TestCommand(t)
	req2.Script = "test"
	req2.Result = model.TestResult + "2"
	id, err := storage.SaveRunCommand(req2)
	assert.NoError(t, err)
	assert.NotNil(t, req2.ID)
	assert.Equal(t, req.ID+1, req2.ID)

	err = storage.DeleteCommand(id)
	assert.NoError(t, err)

	resp, err := storage.GetOneCommand(id)
	assert.Nil(t, resp)

	err = storage.DeleteCommand(999)
	assert.Error(t, err)
}
