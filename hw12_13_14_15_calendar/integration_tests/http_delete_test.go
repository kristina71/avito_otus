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

func TestNegativeDeleteCalendar(t *testing.T) {
	serv := PrepareTest()

	t.Run("error not found", func(t *testing.T) {
		apitest.New().
			Handler(serv).
			Delete("/calendar/delete/13").
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
