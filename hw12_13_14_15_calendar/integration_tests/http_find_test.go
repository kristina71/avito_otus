//go:build integration
// +build integration

package integration_tests

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindCalendar(t *testing.T) {
	serv := PrepareTest()

	t.Run("getEventsPerDay empty request", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Get("/calendar/getEventsPerDay").
			Body("").
			Expect(t).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)

				body, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					require.NoError(t, err)
				}(res.Body)

				fmt.Println("response: ", string(body))
				assert.Equal(t, http.StatusBadRequest, res.StatusCode)
				assert.NotEmpty(t, res)
				return nil
			}).
			End()
	})

	t.Run("getEventsPerDay empty request", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Get("/calendar/getEventsPerWeek").
			Body("").
			Expect(t).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)

				body, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					require.NoError(t, err)
				}(res.Body)

				fmt.Println("response: ", string(body))
				assert.Equal(t, http.StatusBadRequest, res.StatusCode)
				assert.NotEmpty(t, res)
				return nil
			}).
			End()
	})

	t.Run("getEventsPerMonth empty request", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Get("/calendar/getEventsPerMonth").
			Body("").
			Expect(t).
			Assert(func(res *http.Response, req *http.Request) error {
				fmt.Println("request: ", req.Body)

				body, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					require.NoError(t, err)
				}(res.Body)

				fmt.Println("response: ", string(body))
				assert.Equal(t, http.StatusBadRequest, res.StatusCode)
				assert.NotEmpty(t, res)
				return nil
			}).
			End()
	})
}
