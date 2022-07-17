package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type returnCode int

const (
	ok returnCode = iota
	err
	errCmdCannotExecute returnCode = iota + 124
	errCmdNotFound
	errInvalidArg returnCode = iota - 2
)

func TestRunCmd(t *testing.T) {
	type tc struct {
		name               string
		command            string
		args               []string
		env                Environment
		expectedReturnCode returnCode
	}

	tests := []tc{
		{
			name:               "Positive test",
			command:            "ls",
			expectedReturnCode: ok,
		},
		{
			name:               "Permission denied",
			command:            "/bin/bash",
			args:               []string{"-c", "/dev/null"},
			expectedReturnCode: errCmdCannotExecute,
		},
		{
			name:               "Command not found",
			command:            "/bin/bash",
			args:               []string{"-c", "invalid_command"},
			expectedReturnCode: errCmdNotFound,
		},
		{
			name:               "Got an error",
			command:            "/bin/bash",
			args:               []string{"-c", "touch /etc"},
			expectedReturnCode: err,
		},
		{
			name:               "One parameter",
			command:            "/bin/bash",
			args:               []string{"-c"},
			expectedReturnCode: errInvalidArg,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := []string{tc.command}
			cmd = append(cmd, tc.args...)

			code := RunCmd(cmd, tc.env)

			assert.Equal(t, int(tc.expectedReturnCode), code)
		})
	}
}
