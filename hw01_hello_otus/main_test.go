package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRevert(t *testing.T) {
	testCases := []struct {
		name           string
		str            string
		expectedResult string
	}{
		{
			name:           "Test case 1",
			str:            "Hello, OTUS!",
			expectedResult: "!SUTO ,olleH",
		},
		{
			name:           "Test case 2",
			str:            "H",
			expectedResult: "H",
		},
		{
			name:           "Test case 3",
			str:            "",
			expectedResult: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.name, func(t *testing.T) {
				reverted := revertString(testCase.str)
				require.Equal(t, testCase.expectedResult, reverted)
			})
	}
}
