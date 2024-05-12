package tests

import (
	"fmt"
	// "encoding/json"
	sl "main/internal/lib/logger"
	"main/internal/model"
	"main/internal/server"

	// "main/internal/storage/sqlstore"
	"main/internal/storage/teststorage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_HandlerInccorectGetOneCommand(t *testing.T) {

	config := testNewConfig()

	storage := teststorage.New()
	// storage, teardown := sqlstore.TestDB(t, config.StoragePath)
	// defer teardown(sqlstore.Table)

	var logs = sl.SetupLogger(config.Env)
	s := server.NewServer(config, storage, logs)

	testCase := []struct {
		name         string
		command      interface{}
		metod        string
		expectedCode int
	}{

		{
			name: "incorrec method POST",
			command: model.Command{
				Script: "#!/bin/bash\necho \"Hello, World Test!!!\"",
			},
			metod:        "POST",
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

			rec := httptest.NewRecorder()
			req_test := model.TestCommand(t)
			req_test.ID = 1
			storage.Commands["echo 'Hello, World_HOC!'"] = req_test

			qualy := fmt.Sprintf("%s?id=%d", server.PathFind, req_test.ID)
			req, err := http.NewRequest(tc.metod, qualy, nil)
			if err != nil {
				t.Errorf("Failed to create request: %v", err)
			}

			req.Header.Set("Content-Type", "application/json")

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestServer_HandleGetOneCommand(t *testing.T) {

	storage := teststorage.New()

	// storage, teardown := sqlstore.TestDB(t, config.StoragePath)
	// defer teardown(sqlstore.Table)

	config := testNewConfig()
	var logs = sl.SetupLogger(config.Env)
	s := server.NewServer(config, storage, logs)

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
		{
			name:         "valid",
			id:           "2",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "novalid id",
			id:           "",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "novalid id",
			id:           "l",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req_test := model.TestCommand(t)
			req_test.ID = 1
			storage.Commands["echo 'Hello, World_HOC!'"] = req_test

			qualy := fmt.Sprintf("%s?id=%s", server.PathFind, tc.id)
			req, err := http.NewRequest(http.MethodGet, qualy, nil)
			if err != nil {
				t.Errorf("Failed to create request: %v", err)
			}

			req.Header.Set("Content-Type", "application/json")

			s.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}
