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
