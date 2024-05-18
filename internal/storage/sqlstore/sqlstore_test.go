package sqlstore_test

import (
	"fmt"
	"main/internal/model"
	"main/internal/storage/sqlstore"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var (
	databaseURL string
)

// TestMain is a helper function to setup the test database
func TestMain(m *testing.M) {
	// //fmt.Println("databaseURL:", databaseURL)
	nameDatabase := "restapi_test"
	authBase := ""
	flags := "sslmode=disable"
	fmt.Println("//// ", os.Getenv("DATABASE_HOST"), "  ////")
	if os.Getenv("DATABASE_HOST") == "db" {
		authBase = "user:password"
	}
	databaseURL = fmt.Sprintf("postgres://%s@%s:5432/%s?%s", authBase, os.Getenv("DATABASE_HOST"), nameDatabase, flags)
	databaseURL = "postgres://localhost:5432/restapi_test?sslmode=disable"
	fmt.Println(databaseURL)

	os.Exit(m.Run())
}

// TestCommand_CreateRun tests creating a new command
func TestCommand_CreateRun(t *testing.T) {
	// //fmt.Println("||||||", databaseURL, "||||||")
	storage, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown(sqlstore.Table)

	req := model.TestCommand(t)
	req.Result = model.TestResult
	id, err := storage.SaveRunCommand(req)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	assert.NotNil(t, req)
	assert.Equal(t, model.TestResult, req.Result)
}

// TestCommand_GetOneCommand tests getting a single command by ID
func TestCommand_GetOneCommand(t *testing.T) {
	storage, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown(sqlstore.Table)

	req := model.TestCommand(t)
	req.Result = model.TestResult
	id, err := storage.SaveRunCommand(req)
	assert.NoError(t, err)
	assert.NotNil(t, id)

	resp, err := storage.GetOneCommand(id)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Result, resp.Result)

	resp, err = storage.GetOneCommand(999)
	assert.Error(t, err)
	assert.Nil(t, resp)

}

// / TestCommand_GetListCommands tests getting a list of commands by ID
func TestCommand_GetListCommands(t *testing.T) {
	storage, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown(sqlstore.Table)

	resp, err := storage.GetListCommands()
	assert.NoError(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, 0, len(resp))

	req := model.TestCommand(t)
	req.Result = model.TestResult
	_, err = storage.SaveRunCommand(req)
	assert.NoError(t, err)
	assert.NotNil(t, req.ID)

	resp, err = storage.GetListCommands()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, len(resp))

	req2 := model.TestCommand(t)
	req2.Script = req.Script + "2"
	req2.Result = model.TestResult + "2"
	_, err = storage.SaveRunCommand(req2)
	assert.NoError(t, err)
	assert.NotNil(t, req.ID)
	assert.Equal(t, req.ID+1, req2.ID)

	resp, err = storage.GetListCommands()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(resp))
}

// TestCommand_DeleteCommand tests deleting a command by ID
func TestCommand_DeleteCommand(t *testing.T) {
	storage, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown(sqlstore.Table)

	err := storage.DeleteCommand(999)

	assert.Error(t, err)

	req := model.TestCommand(t)
	req.Result = model.TestResult
	_, err = storage.SaveRunCommand(req)
	assert.NoError(t, err)
	assert.NotNil(t, req.ID)

	req2 := model.TestCommand(t)
	req2.Script = "test"
	req2.Result = model.TestResult + "2"
	id, err := storage.SaveRunCommand(req2)
	assert.NoError(t, err)
	assert.NotNil(t, req2.ID)
	assert.Equal(t, req.ID+1, req2.ID)
	////fmt.Println(id)

	resp, err := storage.GetOneCommand(id)
	assert.NotNil(t, resp)

	err = storage.DeleteCommand(id)
	assert.NoError(t, err)

	resp, err = storage.GetOneCommand(id)
	assert.Error(t, err)
	assert.Nil(t, resp)

	err = storage.DeleteCommand(999)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
