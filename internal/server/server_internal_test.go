package server /// TODO:  возможно переименноваьт в пакет в app и тестим по факьу само приложение

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/internal/config"
	sl "main/internal/lib/logger"
	"main/internal/model"
	"main/internal/storage/teststorage"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	// "main/internal/storage/sqlstore" //тут тестовый тестор д.б
)

// !!!TODO: Внутрение веши тестируем с internal_test.go
func testNewConfig() *config.Config {
	config := &config.Config{}
	config.HTTPServer.Address = "localhost:8080"
	config.StoragePath = "host=localhost dbname=restapi_test sslmode=disable"
	Timeout, _ := time.ParseDuration("4s")
	IdleTimeout, _ := time.ParseDuration("1m")
	config.Env = "local"
	config.Address = "localhost:8080"
	config.HTTPServer.Timeout = Timeout
	config.HTTPServer.IdleTimeout = IdleTimeout
	return config
}

func TestServer_HandlerCommandCreate(t *testing.T) {

	storage := teststorage.New()
	config := testNewConfig()
	var logs = sl.SetupLogger(config.Env)
	s := NewServer(config, storage, logs)

	testCase := []struct {
		name         string
		command      interface{}
		expectedCode int
	}{
		// {
		// 	name:         "script is empty",
		// 	command:      model.Command{},
		// 	expectedCode: http.StatusUnprocessableEntity,
		// },
		{
			name: "valid",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
				Result: "Hello, World Test!!!",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "command already exists",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
				Result: "Hello, World Test!!!",
			},
			expectedCode: http.StatusConflict,
		},

		{
			name: "script not correct",
			command: model.Command{
				Script: "not valid script",
				Result: "Hello, World Test!!!",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		// {
		// 	name:         "incorrect request method", //сделать отдельный тест
		// 	command:      nil,
		// 	expectedCode: http.StatusMethodNotAllowed,
		// },
		{
			name: "incorrect request", //сделать отдельный тест
			command: map[string]interface{}{
				"novalid":    "not valid script",
				"novalidres": "Hello, World Test!!!",
			},
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.command) // !!! можно убрать err

			if err != nil {
				t.Fatalf("Failed to encode command: %v", err)
				return
			}

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/command/save", b)

			// req.Header.Set("Content-Type", "application/json")

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)

		})
	}

}

func TestServer_HandleGetOneCommand(t *testing.T) {

	storage := teststorage.New()
	config := testNewConfig()
	var logs = sl.SetupLogger(config.Env)
	s := NewServer(config, storage, logs)

	testCase := []struct {
		name         string
		id           string
		expectedCode int
	}{
		{
			name:         "valid",
			id:           "1",
			expectedCode: http.StatusOK,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req_test := model.TestCommand(t)
			req_test.ID = 1
			storage.Commands["echo 'Hello, World_HOC!'"] = req_test

			qualy := fmt.Sprintf("/command/find?id=%d", req_test.ID)
			req, err := http.NewRequest("GET", qualy, nil)
			if err != nil {
				t.Errorf("Failed to create request: %v", err)
			}

			req.Header.Set("Content-Type", "application/json")

			s.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestServer_HandleGetListCommands(t *testing.T) {

	storage := teststorage.New()
	config := testNewConfig()
	var logs = sl.SetupLogger(config.Env)
	s := NewServer(config, storage, logs)

	testCase := []struct {
		name         string
		expectedCode int
	}{
		{
			name:         "valid",
			expectedCode: http.StatusOK,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/commands/all", nil)

			// req.Header.Set("Content-Type", "application/json")

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestServer_HandleDeleteCommand(t *testing.T) {

	storage := teststorage.New()
	config := testNewConfig()
	var logs = sl.SetupLogger(config.Env)
	s := NewServer(config, storage, logs)

	testCase := []struct {
		name         string
		id           string
		expectedCode int
	}{
		{
			name:         "valid",
			id:           "1",
			expectedCode: http.StatusOK,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req_test := model.TestCommand(t)
			req_test.ID = 1

			storage.Commands["echo 'Hello, World_HDC!'"] = req_test

			qualy := fmt.Sprintf("/command/delete?id=%d", req_test.ID)
			req, err := http.NewRequest("DELETE", qualy, nil)
			if err != nil {
				t.Errorf("Failed to create request: %v", err)
			}

			// req.Header.Set("Content-Type", "application/json")

			s.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

// func TestServer_HandlerTest(t *testing.T) {

// 	storage := teststorage.New()
// 	config := testNewConfig()
// 	var logs = sl.SetupLogger(config.Env)
// 	s := NewServer(config, storage, logs)

// 	testCase := []struct {
// 		name         string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "valid",
// 			expectedCode: 207,
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {

// 			rec := httptest.NewRecorder()

// 			req /*, _ */ := httptest.NewRequest(http.MethodGet, "/test", nil)

// 			// if err != nil {
// 			// 	t.Errorf("Failed to create request: %v", err)

// 			// }
// 			fmt.Println(req)
// 			// http.HandleFunc("/commands/all", s.handleGetListCommands(*s.log))

// 			req.Header.Set("Content-Type", "application/json")

// 			s.ServeHTTP(rec, req)
// 			fmt.Println(rec, " OOOOOOO ", req)

// 			assert.Equal(t, tc.expectedCode, rec.Code)
// 		})
// 	}

// }
