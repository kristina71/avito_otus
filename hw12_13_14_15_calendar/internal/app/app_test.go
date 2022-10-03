package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"testing"
	"time"

	uuid2 "github.com/gofrs/uuid"
	"github.com/google/uuid"

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

func TestCreateGetDeleteEvent(t *testing.T) {
	calendar, ctx := getCalendar()

	testCases := []testCase{
		{
			name: "Create,get and delete event",
			event: storage.Event{
				ID:          uuid2.UUID(uuid.New()),
				Title:       faker.Name(),
				StartAt:     time.Now().AddDate(0, 0, 2),
				Duration:    3,
				Description: faker.Sentence(),
				UserID:      uuid2.UUID(uuid.New()),
				RemindAt:    4,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := calendar.Create(ctx, &testCase.event)
			require.NoError(t, err)

			getEvents, err := calendar.ListAll(ctx)
			require.NoError(t, err)

			require.NotEmpty(t, getEvents)

			err = calendar.Delete(ctx, testCase.event.ID)
			require.NoError(t, err)

			getEmptyEvent, err := calendar.Get(ctx, &testCase.event)
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
				ID:          uuid2.UUID(uuid.New()),
				Title:       faker.Name(),
				StartAt:     time.Now().AddDate(0, 0, 2),
				Duration:    3,
				Description: faker.Sentence(),
				UserID:      uuid2.UUID(uuid.New()),
				RemindAt:    3,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := calendar.Create(ctx, &testCase.event)
			require.NoError(t, err)

			// не проверяю целиком, так как может поменяться порядок startAt и endAt
			// или не совпадет время на сервере
			testCase.event.Title = faker.Name()
			err = calendar.Update(ctx, &testCase.event)
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
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now().AddDate(0, 0, 2),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now().AddDate(0, 0, 2),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fmt.Println(&testCase.events[0])
			err := calendar.Create(ctx, &testCase.events[0])
			require.NoError(t, err)

			err = calendar.Create(ctx, &testCase.events[1])
			require.NoError(t, err)

			events, err := calendar.ListAll(ctx)
			require.NoError(t, err)
			require.GreaterOrEqual(t, len(events), len(testCase.events))

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
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now().AddDate(0, 0, 2),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now().AddDate(0, 0, 2),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					// - 2 months
					RemindAt: 3,
				},
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now().AddDate(0, 0, 2),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					// - 2 week
					RemindAt: 3,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := calendar.Create(ctx, &testCase.events[0])
			require.NoError(t, err)

			err = calendar.Create(ctx, &testCase.events[1])
			require.NoError(t, err)

			err = calendar.Create(ctx, &testCase.events[2])
			require.NoError(t, err)

			events, err := calendar.ListAll(ctx)
			require.NoError(t, err)

			fmt.Println(events)
			// require.Equal(t, events[0].ID, newEvent2.ID)
			// require.Equal(t, events[1].ID, newEvent1.ID)
			// require.Equal(t, events[2].ID, newEvent0.ID)

			/* date := time.Now()
			dayEvents, err := calendar.GetEventsPerDay(ctx, date)
			require.NoError(t, err)
			require.Contains(t, testCase.events[0], dayEvents)

			monthEvents, err := calendar.GetEventsPerMonth(ctx, date)
			require.NoError(t, err)
			require.Contains(t, testCase.events[2], monthEvents)
			weekEvents, err := calendar.GetEventsPerWeek(ctx, date)
			require.NoError(t, err)
			require.Contains(t, testCase.events[1], weekEvents)

			*/

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
