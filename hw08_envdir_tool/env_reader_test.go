package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	BAR   string = "bar"
	EMPTY string = " "
	FOO   string = `   foo
with new line`
	HELLO string = `"hello"`
	UNSET string = ""
)

func TestReadDir(t *testing.T) {
	expectedMap := Environment{
		"BAR":   BAR,
		"EMPTY": strings.TrimRight(EMPTY, " \t"),
		"FOO":   FOO,
		"HELLO": HELLO,
		"UNSET": UNSET,
	}

	const dir = "./testdata/env"
	env, err := ReadDir(dir)
	assert.NoError(t, err)

	assert.Len(t, env, 5)
	assert.Equal(t, expectedMap, env)
}

func TestReadValue(t *testing.T) {
	type testCase struct {
		name           string
		in             string
		expectedResult string
	}

	tests := []testCase{
		{
			name: "BAR: get only first string",
			in: `bar
PLEASE IGNORE SECOND LINE
`,
			expectedResult: BAR,
		},
		{
			name:           "FOO: replace NUL char with 'new line'",
			in:             "   foo\x00with new line",
			expectedResult: FOO,
		},
		{
			name:           "EMPTY: empty line",
			in:             " ",
			expectedResult: strings.TrimRight(EMPTY, " \t"),
		},
		{
			name:           "HELLO: get value with quotes",
			in:             `"hello"`,
			expectedResult: HELLO,
		},
		{
			name:           "UNSET: empty line",
			in:             "",
			expectedResult: UNSET,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := bytes.NewBufferString(tc.in)
			actualResult, err := readValue(b)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
