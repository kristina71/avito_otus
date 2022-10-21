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

func TestNegativeUpdateCalendar(t *testing.T) {
	serv := PrepareTest()

	t.Run("error not found", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Put("/calendar/update/13").
			JSON(`{
			"title": "fgdfg",
			"startAt":"2022-10-02T13:00:00Z",
			"endAt":"2022-10-02T13:00:00Z",
			"duration": 10,
			"description": "",
			"remindAt": 0
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
	t.Run("bad json", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Put("/calendar/update/13").
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
