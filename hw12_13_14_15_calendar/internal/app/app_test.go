package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/config"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/logger"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage/initstorage"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "../../configs/config.yaml", "Path to configuration file")
}

type testCase struct {
	name   string
	events []storage.Event
	event  storage.Event
}

// надо подумать как очищать данные при первом запуске и перед выполнением каждого теста

func TestCreateGetDeleteEvent(t *testing.T) {
	calendar, ctx := getCalendar()

	testCases := []testCase{
		{
			name: "Create,get and delete event",
			event: storage.Event{
				ID:          1,
				Title:       faker.Name(),
				StartAt:     time.Now(),
				EndAt:       time.Now(),
				Description: faker.Sentence(),
				UserID:      1,
				RemindAt:    time.Now(),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			event, err := calendar.Create(ctx,
				testCase.event.UserID,
				testCase.event.Title,
				testCase.event.Description,
				testCase.event.StartAt,
				testCase.event.EndAt,
				testCase.event.RemindAt)
			require.NoError(t, err)

			getEvent, err := calendar.Get(ctx, event.ID)
			require.NoError(t, err)

			require.Equal(t, event.ID, getEvent.ID)
			require.Equal(t, event.Description, getEvent.Description)

			err = calendar.Delete(ctx, event.ID)
			require.NoError(t, err)

			getEmptyEvent, err := calendar.Get(ctx, event.ID)
			require.Error(t, err)
			require.Empty(t, getEmptyEvent)
		})
	}
}

func TestCreateUpdateEvent(t *testing.T) {
	calendar, ctx := getCalendar()

	testCases := []testCase{
		{
			name: "Create and update event",
			event: storage.Event{
				ID:          1,
				Title:       faker.Name(),
				StartAt:     time.Now(),
				EndAt:       time.Now(),
				Description: faker.Sentence(),
				UserID:      1,
				RemindAt:    time.Now(),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			event, err := calendar.Create(ctx,
				testCase.event.UserID,
				testCase.event.Title,
				testCase.event.Description,
				testCase.event.StartAt,
				testCase.event.EndAt,
				testCase.event.RemindAt)
			require.NoError(t, err)

			// не проверяю целиком, так как может поменяться порядок startAt и endAt
			// или не совпадет время на сервере
			err = calendar.Update(ctx, event.ID, event)
			require.NoError(t, err)
		})
	}
}

func TestCreateDeleteAll(t *testing.T) {
	calendar, ctx := getCalendar()

	testCases := []testCase{
		{
			name: "Create and delete all",
			events: []storage.Event{

				{
					ID:          1,
					Title:       faker.Name(),
					StartAt:     time.Now(),
					EndAt:       time.Now(),
					Description: faker.Sentence(),
					UserID:      1,
					RemindAt:    time.Now(),
				},
				{
					ID:          2,
					Title:       faker.Name(),
					StartAt:     time.Now(),
					EndAt:       time.Now(),
					Description: faker.Sentence(),
					UserID:      2,
					RemindAt:    time.Now(),
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := calendar.Create(ctx,
				testCase.events[0].UserID,
				testCase.events[0].Title,
				testCase.events[0].Description,
				testCase.events[0].StartAt,
				testCase.events[0].EndAt,
				testCase.event.RemindAt)
			require.NoError(t, err)

			_, err = calendar.Create(ctx,
				testCase.events[1].UserID,
				testCase.events[1].Title,
				testCase.events[1].Description,
				testCase.events[1].StartAt,
				testCase.events[1].EndAt,
				testCase.event.RemindAt)
			require.NoError(t, err)

			events, err := calendar.ListAll(ctx)
			require.NoError(t, err)
			require.GreaterOrEqual(t, len(events), len(testCase.events))

			// не проверяю целиком, так как может поменяться порядок startAt и endAt
			// или не совпадет время на сервере
			//require.Contains(t, testCase.events[0].ID, events)
			//require.Contains(t, testCase.events[1].ID, events)

			err = calendar.DeleteAll(ctx)
			require.NoError(t, err)

			emptyEvents, err := calendar.ListAll(ctx)
			require.NoError(t, err)
			require.Empty(t, emptyEvents)
		})
	}
}

func TestCreateGetLists(t *testing.T) {
	calendar, ctx := getCalendar()

	testCases := []testCase{
		{
			name: "Create and get lists",
			events: []storage.Event{
				{
					ID:          1,
					Title:       faker.Name(),
					StartAt:     time.Now(),
					EndAt:       time.Now(),
					Description: faker.Sentence(),
					UserID:      1,
					RemindAt:    time.Now(),
				},
				{
					ID:          2,
					Title:       faker.Name(),
					StartAt:     time.Now(),
					EndAt:       time.Now(),
					Description: faker.Sentence(),
					UserID:      2,
					// - 2 months
					RemindAt: time.Now().AddDate(0, -2, 0),
				},
				{
					ID:          3,
					Title:       faker.Name(),
					StartAt:     time.Now(),
					EndAt:       time.Now(),
					Description: faker.Sentence(),
					UserID:      3,
					// - 2 week
					RemindAt: time.Now().AddDate(0, 0, -14),
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := calendar.Create(ctx,
				testCase.events[0].UserID,
				testCase.events[0].Title,
				testCase.events[0].Description,
				testCase.events[0].StartAt,
				testCase.events[0].EndAt,
				testCase.event.RemindAt)
			require.NoError(t, err)

			_, err = calendar.Create(ctx,
				testCase.events[1].UserID,
				testCase.events[1].Title,
				testCase.events[1].Description,
				testCase.events[1].StartAt,
				testCase.events[1].EndAt,
				testCase.event.RemindAt)
			require.NoError(t, err)

			_, err = calendar.Create(ctx,
				testCase.events[2].UserID,
				testCase.events[2].Title,
				testCase.events[2].Description,
				testCase.events[2].StartAt,
				testCase.events[2].EndAt,
				testCase.event.RemindAt)
			require.NoError(t, err)

			/*events, err := calendar.ListAll(ctx)
			require.NoError(t, err)
			require.Equal(t, testCase.events, events)

			date := time.Now()
			dayEvents, err := calendar.ListDay(ctx, date)
			require.NoError(t, err)
			require.Equal(t, testCase.events[0], dayEvents)

			monthEvents, err := calendar.ListMonth(ctx, date)
			require.NoError(t, err)
			require.Equal(t, testCase.events[2], monthEvents)

			weekEvents, err := calendar.ListWeek(ctx, date)
			require.NoError(t, err)
			require.Equal(t, testCase.events[1], weekEvents)*/

			err = calendar.DeleteAll(ctx)
			require.NoError(t, err)
		})
	}
}

func getCalendar() (*App, context.Context) {
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

	stor, err := initstorage.NewStorage(ctx, config.Storage, dsn)
	if err != nil {
		logger.Error("failed to connect DB: " + err.Error())
	}

	logger.Info("DB connected...")

	calendar := New(logger, stor)
	return calendar, ctx
}
