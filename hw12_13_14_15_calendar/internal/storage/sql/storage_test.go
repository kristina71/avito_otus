package sqlstorage

import (
	"context"
	"testing"
	"time"

	uuid2 "github.com/gofrs/uuid"
	"github.com/google/uuid"

	faker "github.com/bxcodec/faker/v3"
	allure "github.com/dailymotion/allure-go"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

type testCase struct {
	name   string
	mock   func(tc *testCase)
	event  storage.Event
	events []storage.Event
}

func TestInsertDB(t *testing.T) {
	allure.Test(t,
		allure.Description("Insert data in DB"),
		allure.Action(func() {
			db, mock, err := sqlxmock.Newx()
			require.NoError(t, err)

			defer db.Close()

			stor := New(db)

			testCases := []testCase{
				{
					name: "OK",
					event: storage.Event{
						ID:          uuid2.UUID(uuid.New()),
						Title:       faker.Name(),
						StartAt:     time.Now(),
						Duration:    3,
						Description: faker.Sentence(),
						UserID:      uuid2.UUID(uuid.New()),
						RemindAt:    3,
					},
					mock: func(tc *testCase) {
						sqlxmock.NewRows([]string{"id"}).AddRow(1)
						mock.ExpectExec("INSERT INTO events").
							WithArgs(tc.event.ID, tc.event.Title, tc.event.StartAt, tc.event.Duration, tc.event.Description,
								tc.event.UserID, tc.event.RemindAt).WillReturnResult(sqlxmock.NewResult(1, 1))
					},
				},
				{
					name: "Insert empty fields",
					event: storage.Event{
						ID:          uuid2.UUID(uuid.New()),
						Title:       "",
						StartAt:     time.Now(),
						Duration:    3,
						Description: "",
						UserID:      uuid2.UUID(uuid.New()),
						RemindAt:    3,
					},
					mock: func(tc *testCase) {
						sqlxmock.NewRows([]string{"id"}).AddRow(1)
						mock.ExpectExec("INSERT INTO events").
							WithArgs(tc.event.ID, tc.event.Title, tc.event.StartAt, tc.event.Duration, tc.event.Description,
								tc.event.UserID, tc.event.RemindAt).WillReturnResult(sqlxmock.NewResult(1, 1))
					},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					allure.Step(allure.Description("Insert data and check result"), allure.Action(func() {
						err := stor.Create(context.Background(), &testCase.event)

						require.NoError(t, err)
					}))
				})
			}
		}))
}

func TestSelectDB(t *testing.T) {
	allure.Test(t,
		allure.Description("Select data in DB"),
		allure.Action(func() {
			db, mock, err := sqlxmock.Newx()
			require.NoError(t, err)

			defer db.Close()

			stor := New(db)

			testCases := []testCase{
				{
					name: "OK",
					event: storage.Event{
						ID:          uuid2.UUID(uuid.New()),
						Title:       faker.Name(),
						StartAt:     time.Now(),
						Duration:    3,
						Description: faker.Sentence(),
						UserID:      uuid2.UUID(uuid.New()),
						RemindAt:    3,
					},
					mock: func(tc *testCase) {
						rows := sqlxmock.NewRows([]string{"id", "title", "start_at", "duration", "description", "user_id", "remind_at"}).
							AddRow(
								tc.event.ID,
								tc.event.Title,
								tc.event.StartAt,
								tc.event.Duration,
								tc.event.Description,
								tc.event.UserID,
								tc.event.RemindAt)
						mock.ExpectQuery("^SELECT (.+) FROM events WHERE id = \\$1").
							WithArgs(tc.event.ID).
							WillReturnRows(rows)
					},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					allure.Step(allure.Description("Get data by id and check result"), allure.Action(func() {
						id, err := stor.Get(context.Background(), &testCase.event)
						require.NoError(t, err)

						require.Equal(t, testCase.event.ID, id)
					}))
				})
			}
		}))
}

func TestDeleteDB(t *testing.T) {
	allure.Test(t,
		allure.Description("Delete data in DB"),
		allure.Action(func() {
			db, mock, err := sqlxmock.Newx()
			require.NoError(t, err)

			defer db.Close()

			stor := New(db)

			testCases := []testCase{
				{
					name: "OK",
					event: storage.Event{
						ID: uuid2.UUID(uuid.New()),
					},
					mock: func(tc *testCase) {
						mock.ExpectExec("^DELETE FROM events WHERE id = \\$1").
							WithArgs(tc.event.ID).WillReturnResult(sqlxmock.NewResult(1, 1))
					},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					allure.Step(allure.Description("Delete data and check result"), allure.Action(func() {
						err = stor.Delete(context.Background(), testCase.event.ID)
						require.NoError(t, err)
					}))
				})
			}
		}))
}

func TestDeleteAllDB(t *testing.T) {
	allure.Test(t,
		allure.Description("Delete all data in DB"),
		allure.Action(func() {
			db, mock, err := sqlxmock.Newx()
			require.NoError(t, err)

			defer db.Close()

			stor := New(db)

			testCases := []testCase{
				{
					name: "OK",
					event: storage.Event{
						ID: uuid2.UUID(uuid.New()),
					},
					mock: func(tc *testCase) {
						mock.ExpectExec("^DELETE FROM events").
							WillReturnResult(sqlxmock.NewResult(1, 1))
					},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					allure.Step(allure.Description("Delete all data and check result"), allure.Action(func() {
						err = stor.DeleteAll(context.Background())
						require.NoError(t, err)
					}))
				})
			}
		}))
}

func TestListAllDB(t *testing.T) {
	allure.Test(t,
		allure.Description("List all data in DB"),
		allure.Action(func() {
			db, mock, err := sqlxmock.Newx()
			require.NoError(t, err)

			defer db.Close()

			stor := New(db)

			testCases := []testCase{
				{
					name: "OK",
					events: []storage.Event{
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
					mock: func(tc *testCase) {
						rows := sqlxmock.NewRows([]string{"id", "title", "start_at", "duration", "description", "user_id", "remind_at"}).
							AddRow(
								tc.event.ID,
								tc.event.Title,
								tc.event.StartAt,
								tc.event.Duration,
								tc.event.Description,
								tc.event.UserID,
								tc.event.RemindAt)
						mock.ExpectQuery("^SELECT id, title, start_at, duration, description, user_id, remind_at FROM events").
							WillReturnRows(rows)
					},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					allure.Step(allure.Description("Select all data and check result"), allure.Action(func() {
						events, err := stor.ListAll(context.Background())
						require.NoError(t, err)
						require.NotEmpty(t, events)
					}))
				})
			}
		}))
}

func TestListDayDB(t *testing.T) {
	t.Skip("need fixing")

	allure.Test(t,
		allure.Description("List day data in DB"),
		allure.Action(func() {
			db, mock, err := sqlxmock.Newx()
			require.NoError(t, err)

			defer db.Close()

			stor := New(db)

			testCases := []testCase{
				{
					name: "OK",
					events: []storage.Event{
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
					mock: func(tc *testCase) {
						rows := sqlxmock.NewRows([]string{"id", "title", "start_at", "duration", "description", "user_id", "remind_at"}).
							AddRow(
								tc.event.ID,
								tc.event.Title,
								tc.event.StartAt,
								tc.event.Duration,
								tc.event.Description,
								tc.event.UserID,
								tc.event.RemindAt)
						mock.ExpectQuery("^SELECT id, title, start_at, duration, description," +
							" user_id, remind_at FROM events WHERE start_at BETWEEN $1 AND $1 + (interval '1d')").
							WillReturnRows(rows)
					},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					allure.Step(allure.Description("Select all data and check result"), allure.Action(func() {
						events, err := stor.ListAll(context.Background())
						require.NoError(t, err)
						require.Equal(t, testCase.events, events)
					}))
				})
			}
		}))
}

func TestListWeekDB(t *testing.T) {
	t.Skip("need fixing")

	allure.Test(t,
		allure.Description("List week data in DB"),
		allure.Action(func() {
			db, mock, err := sqlxmock.Newx()
			require.NoError(t, err)

			defer db.Close()

			stor := New(db)

			testCases := []testCase{
				{
					name: "OK",
					events: []storage.Event{
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
					mock: func(tc *testCase) {
						rows := sqlxmock.NewRows([]string{"id", "title", "start_at", "duration", "description", "user_id", "remind_at"}).
							AddRow(
								tc.event.ID,
								tc.event.Title,
								tc.event.StartAt,
								tc.event.Duration,
								tc.event.Description,
								tc.event.UserID,
								tc.event.RemindAt)
						mock.ExpectQuery("^SELECT id, title, start_at, duration, description, user_id," +
							" remind_at FROM events WHERE start_at BETWEEN $1 AND $1 + (interval '7d')").
							WillReturnRows(rows)
					},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					day := time.Now()
					allure.Step(allure.Description("Select all data and check result"), allure.Action(func() {
						events, err := stor.GetEventsPerWeek(context.Background(), day)
						require.NoError(t, err)
						require.Equal(t, testCase.events, events)
					}))
				})
			}
		}))
}

func TestListMonthDB(t *testing.T) {
	t.Skip("need fixing")

	allure.Test(t,
		allure.Description("List month data in DB"),
		allure.Action(func() {
			db, mock, err := sqlxmock.Newx()
			require.NoError(t, err)

			defer db.Close()

			stor := New(db)

			testCases := []testCase{
				{
					name: "OK",
					events: []storage.Event{
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
					mock: func(tc *testCase) {
						rows := sqlxmock.NewRows([]string{"id", "title", "start_at", "duration", "description", "user_id", "remind_at"}).
							AddRow(
								tc.event.ID,
								tc.event.Title,
								tc.event.StartAt,
								tc.event.Duration,
								tc.event.Description,
								tc.event.UserID,
								tc.event.RemindAt)
						mock.ExpectQuery("^SELECT id, title, start_at, duration, description, " +
							" user_id, remind_at FROM events WHERE start_at BETWEEN $1 AND $1 + (interval '1months')").
							WillReturnRows(rows)
					},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					day := time.Now()
					allure.Step(allure.Description("Select all data and check result"), allure.Action(func() {
						events, err := stor.GetEventsPerMonth(context.Background(), day)
						require.NoError(t, err)
						require.Equal(t, testCase.events, events)
					}))
				})
			}
		}))
}

func TestUpdateDB(t *testing.T) {
	allure.Test(t,
		allure.Description("Update data in DB"),
		allure.Action(func() {
			db, mock, err := sqlxmock.Newx()
			require.NoError(t, err)

			defer db.Close()

			stor := New(db)

			testCases := []testCase{
				{
					name: "OK",
					event: storage.Event{
						ID:          uuid2.UUID(uuid.New()),
						Title:       faker.Name(),
						StartAt:     time.Now(),
						Duration:    3,
						Description: faker.Sentence(),
						RemindAt:    3,
					},
					mock: func(tc *testCase) {
						mock.ExpectExec("^UPDATE events SET title = \\$1, start_at = \\$2, duration = \\$3,"+
							" description = \\$4, remind_at = \\$5 WHERE id = \\$6").
							WithArgs(
								tc.event.Title,
								tc.event.StartAt,
								tc.event.Duration,
								tc.event.Description,
								tc.event.RemindAt,
								tc.event.ID,
							).WillReturnResult(sqlxmock.NewResult(1, 1))
					},
				},
				{
					name: "Update with empty fields",
					event: storage.Event{
						ID:          uuid2.UUID(uuid.New()),
						Title:       "",
						StartAt:     time.Now(),
						Duration:    3,
						Description: "",
						RemindAt:    3,
					},
					mock: func(tc *testCase) {
						mock.ExpectExec("^UPDATE events SET title = \\$1, start_at = \\$2, duration = \\$3,"+
							" description = \\$4, remind_at = \\$5 WHERE id = \\$6").
							WithArgs(tc.event.Title,
								tc.event.StartAt,
								tc.event.Duration,
								tc.event.Description,
								tc.event.RemindAt,
								tc.event.ID,
							).WillReturnResult(sqlxmock.NewResult(1, 1))
					},
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					allure.Step(allure.Description("Update data and check result"), allure.Action(func() {
						err = stor.Update(context.Background(), &testCase.event)
						require.NoError(t, err)
					}))
				})
			}
		}))
}

func mockData(testCase testCase) {
	allure.Step(allure.Description("Mock data"), allure.Action(func() {
		testCase.mock(&testCase)
	}))
}
