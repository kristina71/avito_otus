package sqlstorage

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/dailymotion/allure-go"
	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

type testCase struct {
	name    string
	mock    func(tc *testCase)
	event   storage.Event
	id      uint16
	wantErr bool
}

func TestInsertDB(t *testing.T) {
	allure.Test(t,
		allure.Description("Insert data in DB"),
		allure.Action(func() {
			db, mock, err := sqlxmock.Newx()
			require.NoError(t, err)

			defer db.Close()

			//stor := New(db)

			testCases := []testCase{
				{
					name: "OK",
					event: storage.Event{
						Title:       faker.Name(),
						StartAt:     time.Now(),
						EndAt:       time.Now(),
						Description: faker.Sentence(),
						UserID:      1,
						RemindAt:    time.Now(),
					},
					mock: func(tc *testCase) {
						rows := sqlxmock.NewRows([]string{"id"}).AddRow(1)
						mock.ExpectQuery("INSERT INTO events").WithArgs(tc.event.Title, tc.event.StartAt, tc.event.EndAt, tc.event.Description, tc.event.UserID, tc.event.RemindAt).WillReturnRows(rows)
					},
					id:      1,
					wantErr: false,
				},
				{
					name: "Insert empty fields",
					event: storage.Event{
						Title:       "",
						StartAt:     time.Now(),
						EndAt:       time.Now(),
						Description: "",
						UserID:      1,
						RemindAt:    time.Now(),
					},
					mock: func(tc *testCase) {
						rows := sqlxmock.NewRows([]string{"id"}).AddRow(1)
						mock.ExpectQuery("INSERT INTO events").WithArgs(tc.event.Title, tc.event.StartAt, tc.event.EndAt, tc.event.Description, tc.event.UserID, tc.event.RemindAt).WillReturnRows(rows)
					},
					wantErr: true,
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					allure.Step(allure.Description("Insert data and check result"), allure.Action(func() {
						//insertId, err := eventStorage.Create(context.TODO(), testCase.event)

						require.NoError(t, err)
						if testCase.wantErr != true {
							//require.Equal(t, testCase.id, insertId)
						}
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

			//stor := adapters.New(db)

			testCases := []testCase{
				{
					name: "OK",
					event: storage.Event{
						ID:          string(1),
						Title:       faker.Name(),
						StartAt:     time.Now(),
						EndAt:       time.Now(),
						Description: faker.Sentence(),
						UserID:      1,
						RemindAt:    time.Now(),
					},
					mock: func(tc *testCase) {
						rows := sqlxmock.NewRows([]string{"id", "title", "start_at", "end_at", "descr", "user_id", "remind_at"}).
							AddRow(
								tc.event.ID, tc.event.Title, tc.event.StartAt, tc.event.EndAt, tc.event.Description, tc.event.UserID, tc.event.RemindAt)
						mock.ExpectQuery("^SELECT (.+) FROM events WHERE title = \\$1").
							WithArgs(tc.event.Title).
							WillReturnRows(rows)
					},
					id:      1,
					wantErr: false,
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					allure.Step(allure.Description("Get data by small url and check result"), allure.Action(func() {
						var events = []storage.Event{}
						events, err = memorystorage.New().ListAll(context.Background())

						require.NoError(t, err)
						require.Equal(t, testCase.event, events[0])
						require.Equal(t, len(events), 1)
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

			memorystorage := memorystorage.New()
			//stor := New(db)

			testCases := []testCase{
				{
					name: "OK",
					event: storage.Event{
						ID: string(1),
					},
					mock: func(tc *testCase) {
						mock.ExpectExec("^DELETE FROM events WHERE id = \\$1").
							WithArgs(tc.event.ID).WillReturnResult(sqlxmock.NewResult(1, 1))

					},
					id:      1,
					wantErr: false,
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					allure.Step(allure.Description("Delete data and check result"), allure.Action(func() {
						err = memorystorage.Delete(context.Background(), int(testCase.id))
						require.NoError(t, err)
					}))
				})
			}
		}))
}

func TestUpdateDB(t *testing.T) {
	allure.Test(t,
		allure.Description("Delete data in DB"),
		allure.Action(func() {
			db, mock, err := sqlxmock.Newx()
			require.NoError(t, err)

			defer db.Close()

			//storage := adapters.New(db)
			memorystorage := memorystorage.New()

			testCases := []testCase{
				{
					name: "OK",
					event: storage.Event{
						ID:          string(1),
						Title:       faker.Name(),
						StartAt:     time.Now(),
						EndAt:       time.Now(),
						Description: faker.Sentence(),
						RemindAt:    time.Now(),
					},
					mock: func(tc *testCase) {
						mock.ExpectExec("^UPDATE events SET title = \\$2, start_at = \\$3, end_at = \\$4, description = \\$5, remind_at = \\$6 WHERE id = \\$1").
							WithArgs(tc.event.Title,
								tc.event.StartAt,
								tc.event.EndAt,
								tc.event.Description,
								tc.event.RemindAt,
								tc.event.ID,
							).WillReturnResult(sqlxmock.NewResult(1, 1))
					},
					id:      1,
					wantErr: false,
				},
				{
					name: "Update with empty fields",
					event: storage.Event{
						ID:          string(1),
						Title:       "",
						StartAt:     time.Now(),
						EndAt:       time.Now(),
						Description: "",
						RemindAt:    time.Now(),
					},
					mock: func(tc *testCase) {
						mock.ExpectExec("^UPDATE events SET title = \\$2, start_at = \\$3, end_at = \\$4, description = \\$5, remind_at = \\$6 WHERE id = \\$1").
							WithArgs(tc.event.Title,
								tc.event.StartAt,
								tc.event.EndAt,
								tc.event.Description,
								tc.event.RemindAt,
								tc.event.ID,
							).WillReturnResult(sqlxmock.NewResult(1, 1))
					},
					id:      1,
					wantErr: false,
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					mockData(testCase)

					allure.Step(allure.Description("Update data and check result"), allure.Action(func() {
						err = memorystorage.Update(context.Background(), int(testCase.id), testCase.event)
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
