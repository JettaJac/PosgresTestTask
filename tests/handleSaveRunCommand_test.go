package tests

//Тестим сервер

import (
	"bytes"
	"encoding/json"
	"main/internal/lib/logger"
	"main/internal/lib/slogdiscard"
	"main/internal/model"
	"main/internal/server"

	// "main/internal/storage/sqlstore"
	"main/internal/storage/teststore"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestServer_HandlerInccorectMetodsCreateCommand test the correct method
func TestServer_HandlerInccorectMetodsCreateCommand(t *testing.T) {

	config := testNewConfig()

	storage := teststore.New()
	// storage, teardown := sqlstore.TestDB(t, config.DatabaseURL)
	// defer teardown(sqlstore.Table)

	// var logs = sl.SetupLogger(config.Env)
	var logs = slogdiscard.NewDiscardLogger()
	s := server.NewServer(config, storage, logs)

	testCase := []struct {
		name         string
		command      interface{}
		metod        string
		expectedCode int
	}{

		{
			name: "incorrec method GET",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			metod:        "GET",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name: "incorrec method OPTIONS",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			metod:        "OPTIONS",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name: "incorrec method HEAD",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			metod:        "HEAD",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name: "incorrec method PUT",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			metod:        "PUT",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name: "incorrec method DELETE",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			metod:        "DELETE",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name: "incorrec method CONNECT",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			metod:        "CONNECT",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name: "incorrec method TRACE",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			metod:        "TRACE",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name: "incorrec method PATCH",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			metod:        "PATCH",
			expectedCode: http.StatusMethodNotAllowed,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.command)

			if err != nil {
				t.Fatalf("Failed to encode command: %v", err)
				return
			}

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(tc.metod, server.PathSave, b)

			req.Header.Set("Content-Type", "application/json")

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)

		})
	}

}

// TestServer_HandlerContextCreate checks the correctness of the entered data
func TestServer_HandlerCommandCreate(t *testing.T) {
	// suppressErrors()
	config := testNewConfig()

	storage := teststore.New()
	// storage, teardown := sqlstore.TestDB(t, config.DatabaseURL)
	// defer teardown(sqlstore.Table)

	// var logs = sl.SetupLogger(config.Env)
	var logs = slogdiscard.NewDiscardLogger()
	s := server.NewServer(config, storage, logs)

	testCase := []struct {
		name         string
		command      interface{}
		expectedCode int
	}{
		{
			name:         "script is empty",
			command:      model.Command{},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "valid",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "command already exists",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			expectedCode: http.StatusBadRequest,
		},

		{
			name: "script not correct",
			command: model.Command{
				Script: "not valid script", // !!! можно еще сделать общую проверку на то что выдает запрос, но в целом такая проврека есть
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "incorrect request", //сделать отдельный тест
			command: map[string]interface{}{
				"novalid": "not valid script",
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
			req := httptest.NewRequest(http.MethodPost, server.PathSave, b)

			req.Header.Set("Content-Type", "application/json")

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)

		})
	}

}

// // TestServer_HandlerContextCreate check when submitting an empty request
// func TestServer_HandlerCommandCreate_Empty(t *testing.T) {

// 	config := testNewConfig()
// 	storage := teststore.New()
// 	// storage, teardown := sqlstore.TestDB(t, config.DatabaseURL)
// 	// defer teardown(sqlstore.Table)

// 	var logs = sl.SetupLogger(config.Env)
// 	// var logs = slogdiscard.NewDiscardLogger()
// 	s := server.NewServer(config, storage, logs)

// 	testCase := []struct {
// 		name         string
// 		command      interface{}
// 		expectedCode int
// 	}{

// 		{
// 			name:         "empty request",
// 			command:      nil,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 	}
// 	for _, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {

// 			b := &bytes.Buffer{}
// 			err := json.NewEncoder(b).Encode(tc.command)

// 			if err != nil {
// 				t.Fatalf("Failed to encode command: %v", err)
// 				return
// 			}

// 			rec := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodPost, server.PathSave, b)

// 			req.Header.Set("Content-Type", "application/json")

// 			s.ServeHTTP(rec, req)
// 			// assert.Equal(t, tc.expectedCode, rec.Code)

// 		})
// 	}

// }
