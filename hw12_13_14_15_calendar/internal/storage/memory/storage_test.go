package memorystorage

import (
	"context"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
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
			repo := &mocks.Storage{}
			memorystorage := New1(repo)

			expected := testCase.expectedEvent

			repo.On("Create", context.Background(), expected).Return(expected, nil)

			actualEvent, err := memorystorage.repo.Create(context.Background(), testCase.expectedEvent)
			require.NoError(t, err)

			require.Equal(t, testCase.expectedEvent, actualEvent)
		})
	}
}

func TestGetEvent(t *testing.T) {
	testCases := []testCase{
		{
			name: "Get event by id",
			expectedEvent: storage.Event{
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
			repo := &mocks.Storage{}
			memorystorage := New1(repo)

			repo.On("Get", context.Background(), testCase.expectedEvent.ID).Return(testCase.expectedEvent, nil)

			res, err := memorystorage.repo.Get(context.Background(), testCase.expectedEvent.ID)

			require.NoError(t, err)

			require.Equal(t, testCase.expectedEvent, res)
		})
	}
}

func TestUpdateUrl(t *testing.T) {
	testCases := []testCase{
		{
			name: "Update url",
			expectedEvent: storage.Event{
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
			repo := &mocks.Storage{}
			memorystorage := New1(repo)
			expected := testCase.expectedEvent

			repo.On("Update", context.Background(), testCase.expectedEvent.ID, expected).Return(nil)

			err := memorystorage.repo.Update(context.Background(), testCase.expectedEvent.ID, testCase.expectedEvent)
			require.NoError(t, err)
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []testCase{
		{
			name: "Delete url",
			expectedEvent: storage.Event{
				ID: 1,
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
				ID: 1,
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
				{
					ID:          3,
					Title:       faker.Name(),
					StartAt:     time.Now(),
					EndAt:       time.Now(),
					Description: faker.Sentence(),
					UserID:      3,
					RemindAt:    time.Now(),
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
				{
					ID:          3,
					Title:       faker.Name(),
					StartAt:     time.Now(),
					EndAt:       time.Now(),
					Description: faker.Sentence(),
					UserID:      3,
					RemindAt:    time.Now(),
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
			repo.On("ListDay", context.Background(), day).Return(testCase.expectedEvents, nil)

			events, err := memorystorage.repo.ListDay(context.Background(), day)
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
				{
					ID:          3,
					Title:       faker.Name(),
					StartAt:     time.Now(),
					EndAt:       time.Now(),
					Description: faker.Sentence(),
					UserID:      3,
					RemindAt:    time.Now(),
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
			repo.On("ListWeek", context.Background(), day).Return(testCase.expectedEvents, nil)

			events, err := memorystorage.repo.ListWeek(context.Background(), day)
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
				{
					ID:          3,
					Title:       faker.Name(),
					StartAt:     time.Now(),
					EndAt:       time.Now(),
					Description: faker.Sentence(),
					UserID:      3,
					RemindAt:    time.Now(),
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
			repo.On("ListMonth", context.Background(), day).Return(testCase.expectedEvents, nil)

			events, err := memorystorage.repo.ListMonth(context.Background(), day)
			require.NoError(t, err)
			require.Equal(t, testCase.expectedEvents, events)
		})
	}
}
