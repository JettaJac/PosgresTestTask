package server /// TODO:  возможно переименноваьт в пакет в app и тестим по факьу само приложение

import (
	"bytes"
	// "encoding/json"
	// "github.com/stretchr/testify/assert"
	"main/internal/config"
	"main/internal/model"
	"main/internal/storage/sqlstore"
	"main/internal/storage/teststorage"
	"net/http"
	"net/http/httptest"
	"testing"
	// "time"
	"fmt"
	"io"
	"log"
	"main/internal/lib/logger"
	// "main/internal/storage/sqlstore" //тут тестовый тестор д.б
)

// !!!TODO: Внутрение веши тестируем с internal_test.go
func testNewConfig() *config.Config {
	return &config.Config{
		// Env: "local",
		// config.Timeout, _ = time.ParseDuration("4s")
		// IdleTimeout, _ := time.ParseDuration("1m")
	}
}

func TestServer_HandlerCommandCreate(t *testing.T) {
	// // config := config.NewConfig()
	// // var  config *config.Config

	storage := teststorage.New()

	// // fmt.Println(storage)
	// // config := &config.Config{}

	config := testNewConfig()
	storage2, _ := sqlstore.NewDB(config.StoragePath)
	_ = storage2
	_ = storage
	// Timeout, _ := time.ParseDuration("4s")
	// IdleTimeout, _ := time.ParseDuration("1m")

	// config.Env = "local"
	// //config.StoragePath = "./storage/storage.db"
	// config.DatabaseURL = "host=localhost dbname=restapi_test sslmode=disable"
	// config.Address = "localhost:8080"
	// config.Timeout = Timeout
	// config.IdleTimeout = IdleTimeout
	var logs = sl.SetupLogger(config.Env)
	s := NewServer(config, storage2, logs)

	testCase := []struct {
		name         string
		command      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			command: map[string]interface{}{
				"script": "#!/bin/bash\necho \"Hello, World\"",
				"result": "Hello, World",
			},
			expectedCode: http.StatusCreated,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			b := &bytes.Buffer{}
			// err := json.NewEncoder(b).Encode(tc.command) // !!! можно убрать err

			// if err != nil {
			// 	t.Fatalf("Failed to encode command: %v", err)
			// 	return
			// }
			body := []byte(`{"script":"echo 'Hello, World_bytepp!'"}`)
			_ = body
			// _ = b

			rec := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/command/save", b)
			req2, err := http.NewRequest("POST" /*http.MethodPost*/, "http://localhost:8080/command/save", bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
				return
			}

			client := &http.Client{}

			// Отправка запроса на сервер
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			var bb bytes.Buffer
			// Чтение и обработка ответа
			_, err = io.Copy(&bb, resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println()
			fmt.Println(string(bb.Bytes()))
			fmt.Println()

			fmt.Println(resp)
			fmt.Println()

			fmt.Println("UUUUUUUUU   ", resp, "ЛЛЛЛЛЛЛЛ   ", req)
			fmt.Println()
			storage.Commands["iiii"] = &model.Command{Script: "echo 'PPPPPPPPPPP!'"}
			s.router.ServeHTTP(rec, req2)
			// _ = req2

			fmt.Println("UUUUUUUUU   ", rec, "ЛЛЛЛЛЛЛЛ   ", req2, " OOOOOOO ", rec.Code, "__", rec.Body, "---")
			fmt.Println("YYYYYY   ", storage)

			// assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
