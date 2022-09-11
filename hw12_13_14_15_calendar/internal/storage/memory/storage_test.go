package memorystorage

import (
	"context"
	"testing"
	"time"

	uuid2 "github.com/gofrs/uuid"
	"github.com/google/uuid"

	faker "github.com/bxcodec/faker/v3"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage/mocks"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name           string
	expectedEvents []storage.Event
	expectedEvent  storage.Event
}

func TestCreate(t *testing.T) {
	testCases := []testCase{
		{
			name: "Create event",
			expectedEvent: storage.Event{
				ID:          uuid2.UUID(uuid.New()),
				Title:       faker.Name(),
				StartAt:     time.Now(),
				Duration:    3,
				Description: faker.Sentence(),
				UserID:      uuid2.UUID(uuid.New()),
				RemindAt:    3,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := &mocks.Storage{}
			memorystorage := New1(repo)

			repo.On("Create", context.Background(), &testCase.expectedEvent).Return(nil)

			err := memorystorage.repo.Create(context.Background(), &testCase.expectedEvent)
			require.NoError(t, err)
		})
	}
}

func TestGetEvent(t *testing.T) {
	testCases := []testCase{
		{
			name: "Get event by id",
			expectedEvent: storage.Event{
				ID:          uuid2.UUID(uuid.New()),
				Title:       faker.Name(),
				StartAt:     time.Now(),
				Duration:    3,
				Description: faker.Sentence(),
				UserID:      uuid2.UUID(uuid.New()),
				RemindAt:    3,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := &mocks.Storage{}
			memorystorage := New1(repo)

			repo.On("Get", context.Background(), &testCase.expectedEvent).Return(testCase.expectedEvent.ID, nil)

			res, err := memorystorage.repo.Get(context.Background(), &testCase.expectedEvent)
			require.NoError(t, err)

			require.Equal(t, testCase.expectedEvent.ID, res)
		})
	}
}

func TestUpdate(t *testing.T) {
	testCases := []testCase{
		{
			name: "Update event",
			expectedEvent: storage.Event{
				ID:          uuid2.UUID(uuid.New()),
				Title:       faker.Name(),
				StartAt:     time.Now(),
				Duration:    3,
				Description: faker.Sentence(),
				UserID:      uuid2.UUID(uuid.New()),
				RemindAt:    3,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := &mocks.Storage{}
			memorystorage := New1(repo)
			expected := testCase.expectedEvent

			repo.On("Update", context.Background(), &expected).Return(nil)

			err := memorystorage.repo.Update(context.Background(), &testCase.expectedEvent)
			require.NoError(t, err)
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []testCase{
		{
			name: "Delete url",
			expectedEvent: storage.Event{
				ID: uuid2.UUID(uuid.New()),
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := &mocks.Storage{}
			memorystorage := New1(repo)

			repo.On("Delete", context.Background(), testCase.expectedEvent.ID).Return(nil)

			err := memorystorage.repo.Delete(context.Background(), testCase.expectedEvent.ID)
			require.NoError(t, err)
		})
	}
}

func TestDeleteAll(t *testing.T) {
	testCases := []testCase{
		{
			name: "Delete all",
			expectedEvent: storage.Event{
				ID: uuid2.UUID(uuid.New()),
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := &mocks.Storage{}
			memorystorage := New1(repo)

			repo.On("DeleteAll", context.Background()).Return(nil)

			err := memorystorage.repo.DeleteAll(context.Background())
			require.NoError(t, err)
		})
	}
}

func TestListAll(t *testing.T) {
	testCases := []testCase{
		{
			name: "List all",
			expectedEvents: []storage.Event{
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
			},
		},
		{
			name: "Empty list",
			expectedEvents: []storage.Event{
				{},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := &mocks.Storage{}
			memorystorage := New1(repo)

			repo.On("ListAll", context.Background()).Return(testCase.expectedEvents, nil)

			events, err := memorystorage.repo.ListAll(context.Background())
			require.NoError(t, err)
			require.Equal(t, testCase.expectedEvents, events)
		})
	}
}

func TestListDay(t *testing.T) {
	testCases := []testCase{
		{
			name: "List all",
			expectedEvents: []storage.Event{
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
			},
		},
		{
			name: "Empty list",
			expectedEvents: []storage.Event{
				{},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := &mocks.Storage{}
			memorystorage := New1(repo)
			day := time.Now()
			repo.On("GetEventsPerDay", context.Background(), day).Return(testCase.expectedEvents, nil)

			events, err := memorystorage.repo.GetEventsPerDay(context.Background(), day)
			require.NoError(t, err)
			require.Equal(t, testCase.expectedEvents, events)
		})
	}
}

func TestListWeek(t *testing.T) {
	testCases := []testCase{
		{
			name: "List week",
			expectedEvents: []storage.Event{
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
			},
		},
		{
			name: "Empty list week",
			expectedEvents: []storage.Event{
				{},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := &mocks.Storage{}
			memorystorage := New1(repo)
			day := time.Now()
			repo.On("GetEventsPerWeek", context.Background(), day).Return(testCase.expectedEvents, nil)

			events, err := memorystorage.repo.GetEventsPerWeek(context.Background(), day)
			require.NoError(t, err)
			require.Equal(t, testCase.expectedEvents, events)
		})
	}
}

func TestListMonth(t *testing.T) {
	testCases := []testCase{
		{
			name: "List month",
			expectedEvents: []storage.Event{
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
				{
					ID:          uuid2.UUID(uuid.New()),
					Title:       faker.Name(),
					StartAt:     time.Now(),
					Duration:    3,
					Description: faker.Sentence(),
					UserID:      uuid2.UUID(uuid.New()),
					RemindAt:    3,
				},
			},
		},
		{
			name: "Empty list month",
			expectedEvents: []storage.Event{
				{},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := &mocks.Storage{}
			memorystorage := New1(repo)
			day := time.Now()
			repo.On("GetEventsPerMonth", context.Background(), day).Return(testCase.expectedEvents, nil)

			events, err := memorystorage.repo.GetEventsPerMonth(context.Background(), day)
			require.NoError(t, err)
			require.Equal(t, testCase.expectedEvents, events)
		})
	}
}
