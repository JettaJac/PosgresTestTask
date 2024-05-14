package tests

import (
	"bytes"
	"encoding/json"
	sl "main/internal/lib/logger"
	"main/internal/model"
	"main/internal/server"
	"main/internal/storage/sqlstore"
	"main/internal/storage/teststorage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestServer_HandlerInccorectMetodsGetListCommands tests the correct method
func TestServer_HandlerInccorectMetodsGetListCommands(t *testing.T) {

	config := testNewConfig()

	// storage := teststorage.New()
	storage, teardown := sqlstore.TestDB(t, config.DatabaseURL)
	defer teardown(sqlstore.Table)

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

			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.command)

			if err != nil {
				t.Fatalf("Failed to encode command: %v", err)
				return
			}

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(tc.metod, server.PathList, nil)

			req.Header.Set("Content-Type", "application/json")

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

// TestServer_HandleGetListCommands tests the correct data
func TestServer_HandleGetListCommands(t *testing.T) {

	storage := teststorage.New()
	// storage, teardown := sqlstore.TestDB(t, config.StoragePath)
	// defer teardown(sqlstore.Table)

	config := testNewConfig()
	var logs = sl.SetupLogger(config.Env)
	s := server.NewServer(config, storage, logs)

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
			req := httptest.NewRequest(http.MethodGet, server.PathList, nil)
			req.Header.Set("Content-Type", "application/json")

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}
