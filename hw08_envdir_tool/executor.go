package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	c := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout

	c.Env = SetUpCmdEnv(env)

	if c.Run() != nil {
		returnCode = c.ProcessState.ExitCode()
	}

	return
}

func SetUpCmdEnv(env Environment) []string {
	for k, v := range env {
		if v == "" {
			os.Unsetenv(k)
			continue
		}

		os.Setenv(k, v)
	}

	return os.Environ()
}
