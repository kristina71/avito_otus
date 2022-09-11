package internalhttp

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/config"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage/initstorage"

	"github.com/steinfletcher/apitest"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/app"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/logger"
	"github.com/stretchr/testify/assert"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "../../../configs/config.yaml", "Path to configuration file")
}

func TestAPICalendar(t *testing.T) {
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

	server := NewServer(logger, calendar)

	t.Run("create event", func(t *testing.T) {
		apitest.New().
			HandlerFunc(server.createEvent).
			Post("/calendar/add").
			JSONFromFile("../../../internal/server/http/tests/event.json").
			Expect(t).
			Status(http.StatusOK).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)
				assert.Equal(t, http.StatusOK, res.StatusCode)
				return nil
			}).
			End()

		apitest.New().
			HandlerFunc(server.getEvents).
			Get("/calendar/get").
			Body("").
			Expect(t).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)
				fmt.Println("result: ", res.Body)
				assert.Equal(t, http.StatusOK, res.StatusCode)
				return nil
			}).
			End()

		apitest.New().
			HandlerFunc(server.deleteAllEvents).
			Delete("/calendar/deleteall").
			Expect(t).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)
				fmt.Println("result: ", res.Body)
				assert.Equal(t, http.StatusOK, res.StatusCode)
				return nil
			}).
			End()
	})
}
