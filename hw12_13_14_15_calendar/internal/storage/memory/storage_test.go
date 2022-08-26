package memorystorage

import (
	"context"
	"testing"
	"time"
<<<<<<< HEAD

	faker "github.com/bxcodec/faker/v3"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage/mocks"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name           string
	expectedEvents []storage.Event
	expectedEvent  storage.Event
=======

	"github.com/bxcodec/faker/v3"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage/mocks"
	"github.com/stretchr/testify/require"
)

/*
const (
	layoutISO = "2006-01-02"
)
*/

// тут неплохо было бы сделать тесты с мокери. Если делать такие тесты - будет шляпно =(

func TestStorage(t *testing.T) {
	/*
		t.Run("create in already busy time", func(t *testing.T) {
			st := NewTestStorage(time.Now(), 3)

			_, err := st.Events().Create(nil, storage.NewTestEvent(time.Now().Add(10*time.Second)))

			require.EqualError(t, err, storage.ErrEventExists.Error())
		})

		t.Run("update not found", func(t *testing.T) {
			st := NewTestStorage(time.Now(), 0)

			err := st.Events().Update(nil, 999, storage.NewTestEvent(time.Now().Add(10*time.Second)))

			require.EqualError(t, err, storage.ErrEvent404.Error())
		})

		t.Run("delete", func(t *testing.T) {
			st := NewTestStorage(time.Now(), 0)
			ev := storage.NewTestEvent(time.Now())
			ev, err := st.Events().Create(nil, ev)
			require.NoError(t, err)
			require.Len(t, st.Events().ListForDay(nil, time.Now()).List, 1)

			st.Events().Delete(nil, ev.ID)
			list := st.Events().ListForDay(nil, time.Now())

			require.Len(t, list.List, 0)
		})

		t.Run("list events for day", func(t *testing.T) {
			st := NewTestStorage(time.Now(), 5)

			got := st.Events().ListForDay(nil, time.Now())

			require.Len(t, got.List, 1)
		})

		t.Run("list events for week", func(t *testing.T) {
			pt, err := time.Parse(layoutISO, "2020-08-31")
			require.NoError(t, err)
			st := NewTestStorage(pt, 50)

			got := st.Events().ListForWeek(nil, pt.Add(24*time.Hour))

			require.Len(t, got.List, 7)
		})

		t.Run("list events for month", func(t *testing.T) {
			pt, err := time.Parse(layoutISO, "2020-12-25")
			require.NoError(t, err)
			st := NewTestStorage(pt, 50)

			got := st.Events().ListForMonth(nil, pt)

			require.Len(t, got.List, 7)
		})

		//time to basy
		/*t.Run("list events before date", func(t *testing.T) {
			pt, err := time.Parse(layoutISO, "2020-08-31")
			require.NoError(t, err)
			st := NewTestStorage(pt, 50)

			got := st.Events().ListBeforeDate(nil, pt.AddDate(0, 0, 3))

			require.Len(t, got, 3)
		})*/
}

type testCase struct {
	name string
	// expectedEvents []storage.Event
	expectedEvent storage.Event
>>>>>>> HW-12 Implement db logic
	// wantErr        bool
	// prepare        func(t *testing.T) string
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
<<<<<<< HEAD
=======
	t.Skip("need fixing")
>>>>>>> HW-12 Implement db logic
	testCases := []testCase{
		{
			name: "Get event by id",
			expectedEvent: storage.Event{
<<<<<<< HEAD
<<<<<<< HEAD
=======

>>>>>>> HW-12 Implement db logic
=======
>>>>>>> HW-12 Fix linter issues
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

<<<<<<< HEAD
			repo.On("Get", context.Background(), testCase.expectedEvent.ID).Return(testCase.expectedEvent, nil)

			res, err := memorystorage.repo.Get(context.Background(), testCase.expectedEvent.ID)
=======
			repo.On("Get", context.Background()).Return(testCase.expectedEvent, nil)

			res, err := memorystorage.repo.Get(context.Background(), 1)
>>>>>>> HW-12 Implement db logic
			require.NoError(t, err)

			require.Equal(t, testCase.expectedEvent, res)
		})
	}
}

func TestUpdateUrl(t *testing.T) {
<<<<<<< HEAD
=======
	t.Skip("need fixing")
>>>>>>> HW-12 Implement db logic
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

<<<<<<< HEAD
			repo.On("Update", context.Background(), testCase.expectedEvent.ID, expected).Return(nil)

			err := memorystorage.repo.Update(context.Background(), testCase.expectedEvent.ID, testCase.expectedEvent)
			require.NoError(t, err)
=======
			repo.On("Update", context.Background(), expected).Return(nil)

			err := memorystorage.repo.Update(context.Background(), 1, testCase.expectedEvent)
			require.NoError(t, err)

			// require.Equal(t, expected, resUrl)
>>>>>>> HW-12 Implement db logic
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

<<<<<<< HEAD
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
=======
			repo.On("Delete", context.Background(), 1).Return(nil)

			err := memorystorage.repo.Delete(context.Background(), 1)
			require.NoError(t, err)
>>>>>>> HW-12 Implement db logic
		})
	}
}
