//go:build integration
// +build integration

package integration_tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateCalendar(t *testing.T) {
	serv := PrepareTest()

	t.Run("create event with empty title", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Post("/calendar/add").
			JSON(`{
			"title": "",
			"startAt":"2022-10-02T13:00:00Z",
			"endAt":"2022-10-02T13:00:00Z",
			"duration": 30,
			"description": "1imp",
			"remindAt": 3
	}`).
			Expect(t).
			Status(http.StatusInternalServerError).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)

				_, err := ioutil.ReadAll(res.Body)
				require.Nil(t, err)

				assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
				return nil
			}).
			End()
	})
	t.Run("create event with empty startAt", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Post("/calendar/add").
			JSON(`{
			"title": "dfsdfs",
			"startAt":"",
			"endAt":"2022-10-02T13:00:00Z",
			"duration": 30,
			"description": "1imp",
			"remindAt": 3
	}`).
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)

				_, err := ioutil.ReadAll(res.Body)
				require.Nil(t, err)

				assert.Equal(t, http.StatusBadRequest, res.StatusCode)
				return nil
			}).
			End()
	})
	t.Run("create event with empty endAt", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Post("/calendar/add").
			JSON(`{
			"title": "dfsdfs",
			"startAt":"2022-10-02T13:00:00Z",
			"endAt":"",
			"duration": 30,
			"description": "1imp",
			"remindAt": 3
	}`).
			Expect(t).
			Status(http.StatusOK).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)

				_, err := ioutil.ReadAll(res.Body)
				require.Nil(t, err)
				return nil
			}).
			End()
	})
	t.Run("create event with empty duration", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Post("/calendar/add").
			JSON(`{
			"title": "fgdfg",
			"startAt":"2022-10-02T13:00:00Z",
			"endAt":"2022-10-02T13:00:00Z",
			"duration": 0,
			"description": "1imp",
			"remindAt": 3
	}`).
			Expect(t).
			Status(http.StatusInternalServerError).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)

				_, err := ioutil.ReadAll(res.Body)
				require.Nil(t, err)
				assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
				return nil
			}).
			End()
	})
	t.Run("create event with empty description", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Post("/calendar/add").
			JSON(`{
			"title": "fgdfg",
			"startAt":"2022-09-02T13:00:00Z",
			"endAt":"2022-09-02T13:00:00Z",
			"duration": 10,
			"description": "",
			"remindAt": 3
	}`).
			Expect(t).
			Status(http.StatusOK).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)

				_, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, res.StatusCode)
				return nil
			}).
			End()
	})
	t.Run("create event with empty remindAt", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Post("/calendar/add").
			JSON(`{
			"title": "fgdfg",
			"startAt":"2022-12-02T13:00:00Z",
			"endAt":"2022-12-02T13:00:00Z",
			"duration": 10,
			"description": "",
			"remindAt": 0
	}`).
			Expect(t).
			Status(http.StatusOK).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)

				_, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, res.StatusCode)
				return nil
			}).
			End()
	})
	t.Run("bad json", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Post("/calendar/add").
			JSON(`{
			"title": "fgdfg",
			"startAt":"2022-10-02T13:00:00Z",
			"endAt":"2022-10-02T13:00:00Z",
			"duration": 10,
			"description": "",
			"remindAt": 0
	`).
			Expect(t).
			Status(http.StatusBadRequest).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)

				_, err := ioutil.ReadAll(res.Body)
				require.Nil(t, err)
				assert.Equal(t, http.StatusBadRequest, res.StatusCode)
				return nil
			}).
			End()
	})
}
