package tests

import (
	sl "main/internal/lib/logger"
	"main/internal/server"
	// "main/internal/storage/sqlstore"
	"main/internal/storage/teststorage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			req := httptest.NewRequest(http.MethodGet, "/commands/all", nil)
			req.Header.Set("Content-Type", "application/json")

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}
