package server /// TODO:  возможно переименноваьт в пакет в app и тестим по факьу само приложение

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"main/internal/storage/teststorage"
	"net/http"
	"net/http/httptest"
	"main/internal/config"
	// "main/internal/storage/sqlstore" //тут тестовый тестор д.б
)
// !!!TODO: Внутрение веши тестируем с internal_test.go

func TestServer_HandlerCommandCreate(t *testing.T) {
	config := config.NewConfig()
	storage := teststorage.New()
	s := NewServer(config, storage)
	testCase := []struct {
		name         string
		script     string
		result string
		expectedCode int
	}{
		{
			name: "valid",
			script: "#!/bin/bash\necho \"Hello, World\"",
			result: "Hello, World",		
			expectedCode: http.StatusCreated,	
		},
		

	}
		for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.script)
			req, _ := http.NewRequest("POST", "/command", b)
			s.router.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}