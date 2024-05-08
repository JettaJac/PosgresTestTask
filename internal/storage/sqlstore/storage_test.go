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

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}

func TestCommand_CreateRun(t *testing.T) { // good  возможно тут надо сравнивать результат
	storage, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commandsdb")

	// обработать ошибку
	с := model.TestCommand(t)
	_, err := storage.SaveRunScript(с)
	assert.NoError(t, err)
	assert.NotNil(t, с.ID)
}

func TestCommand_GetOneScript(t *testing.T) { // good  возможно тут надо сравнивать результат
	storage, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commandsdb")

	// обработать ошибку
	c := model.TestCommand(t)
	_, err := storage.SaveRunScript(c)
	assert.NoError(t, err)

	err = storage.GetOneScript(c) // возможно метод должен принимать только id,  а отдавать модель
	assert.NoError(t, err)
	assert.NotNil(t, c)
	// assert.Equal(t, 1, id)
}

func TestCommand_GetListCommands(t *testing.T) { // good  возможно тут надо сравнивать результат
	storage, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("commandsdb")

	// обработать ошибку
	с := model.TestCommand(t)
	_, err := storage.SaveRunScript(с)
	с2 := model.TestCommand(t)
	с2.Name = "test_2" // убрать
	с2.Script = "#!/bin/bash\necho \"Hello, World_Second\""
	_, err = storage.SaveRunScript(с2)
	if err != nil {
		fmt.Println(err)
		// return
	}
	res, err := storage.GetListCommands()
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 2, len(res))
}
