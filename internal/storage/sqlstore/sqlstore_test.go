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
	defer teardown(sqlstore.Table) // !!! возможно прокидываьт сюда название таблицы с которой работаем
	// _ = teardown // !!!

	// обработать ошибку

	req := model.TestCommand(t)
	req.Result = model.TestResult
	id, err := storage.SaveRunScript(req)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	assert.NotNil(t, req)
	assert.Equal(t, model.TestResult, req.Result) // проверяем, что мы не потеряли результат
}

func TestCommand_GetOneScript(t *testing.T) { // good  возможно тут надо сравнивать результат
	storage, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown(sqlstore.Table)
	// _ = teardown // !!!

	req := model.TestCommand(t)
	req.Result = model.TestResult
	id, err := storage.SaveRunScript(req)
	assert.NoError(t, err)
	assert.NotNil(t, id)

	resp, err := storage.GetOneScript(id) // возможно метод должен принимать только id,  а отдавать модель
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Result, resp.Result)

	resp, err = storage.GetOneScript(999)
	assert.Error(t, err)
	assert.Nil(t, resp)

}

// тест на пустой список, !!!
func TestCommand_GetListCommands(t *testing.T) { // good  возможно тут надо сравнивать результат
	storage, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown(sqlstore.Table)

	req := model.TestCommand(t)
	req.Result = model.TestResult
	_, err := storage.SaveRunScript(req)
	assert.NoError(t, err)
	assert.NotNil(t, req.ID)

	req2 := model.TestCommand(t)
	req2.Script = req.Script + "2"
	req2.Result = model.TestResult + "2"
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
	storage, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown(sqlstore.Table)
	// _ = teardown // !!!

	err := storage.DeleteCommand(999)
	// fmt.Println(err)
	assert.Error(t, err)

	req := model.TestCommand(t)
	req.Result = model.TestResult
	_, err = storage.SaveRunScript(req)
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

	resp, err := storage.GetOneScript(id)
	assert.NotNil(t, resp)

	err = storage.DeleteCommand(id)
	assert.NoError(t, err)
	// assert.Nil(t, req2)

	resp, err = storage.GetOneScript(id)
	assert.Nil(t, resp)
	// assert.Equal(t, 2, len(resp))

	err = storage.DeleteCommand(999)
	// fmt.Println(err)
	assert.Error(t, err)

	// assert.Equal(t, err, store.ErrCommandNotFound)
}
