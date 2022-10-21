//go:build integration
// +build integration

package integration_tests

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/app"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/config"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/logger"
	server "github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/server/http"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage/initstorage"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "../configs/config.yaml", "Path to configuration file")
}

func TestE2eCalendar(t *testing.T) {
	serv := PrepareTest()

	apitest.New().
		Handler(serv).
		Post("/calendar/add").
		JSON(`{
			"title": "1test",
			"startAt":"2022-10-02T13:00:00Z",
			"endAt":"2022-10-02T13:00:00Z",
			"duration": 30,
			"description": "1imp",
			"remindAt": 3
	}`).
		Expect(t).
		Status(http.StatusOK).
		Assert(func(res *http.Response, req *http.Request) error {
			fmt.Println("request: ", req.Body)

			body, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				require.NoError(t, err)
			}(res.Body)

			fmt.Println("response: ", string(body))
			assert.Equal(t, http.StatusOK, res.StatusCode)
			return nil
		}).
		End()

	apitest.New().
		Handler(serv).
		Get("/calendar/get").
		Body("").
		Expect(t).
		Assert(func(res *http.Response, req *http.Request) error {
			fmt.Println("request: ", req.Body)

			body, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				require.NoError(t, err)
			}(res.Body)

			fmt.Println("response: ", string(body))
			assert.Equal(t, http.StatusOK, res.StatusCode)
			assert.NotEmpty(t, res)
			return nil
		}).
		End()

	apitest.New().
		Handler(serv).
		Delete("/calendar/deleteall").
		Expect(t).
		Assert(func(res *http.Response, req *http.Request) error {
			fmt.Println("request: ", req.Body)

			body, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				require.NoError(t, err)
			}(res.Body)

			fmt.Println("response: ", string(body))
			assert.Equal(t, http.StatusOK, res.StatusCode)
			assert.Empty(t, err)
			return nil
		}).
		End()
}

func PrepareTest() http.Handler {
	config, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	logger, err := logger.New(config.Logger.Level, config.Logger.File)
	if err != nil {
		log.Fatalf("Logger error: %v", err)
	}

	ctx := context.Background()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.Database.Username,
		config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name,
		config.Database.SSLMode)

	stor, err := initstorage.NewStorage(ctx, "memory", dsn)
	if err != nil {
		logger.Error("failed to connect DB: " + err.Error())
	}

	logger.Info("DB connected...")

	calendar := app.New(logger, stor)
	service := server.NewServer1(logger, calendar)
	serv := server.NewRouter(service)
	return serv
}
