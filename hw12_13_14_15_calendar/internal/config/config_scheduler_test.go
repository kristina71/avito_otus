package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	tempDirectorySchedulerPath = "./"
	tempDirSchedulerName       = "tempdir"
	tempFileSchedulerName      = "tempconfig*.yaml"
	testConfigSchedulerContent = `
logger:
  level: "INFO1"
  path: "log_schedule1.log"
schedule:
  period: "1m"
  remind_for: "1h"
  uri: "amqp://guest:guest@localhost:5672/"
  queue: "calendar"
server:
  host: "127.0.0.1"
  http_port: "8080"
  grpc_port: "50051"
database:
  host: "127.0.0.1"
  port: "5432"
  username: "postgres1"
  password: "password1"
  name: "calendar"
  SSLMode: "disable"
storage: "SQL"`
)

func TestConfigSchedulerReading(t *testing.T) {
	// create temp directory
	tempDir, err := ioutil.TempDir(tempDirectorySchedulerPath, tempDirSchedulerName)
	require.NoErrorf(t, err, "unable to create temp directory")
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			require.NoErrorf(t, err, "unable to delete temp dir %s", tempDir)
		}
	}()
	// crete temp file in temp directory
	tempFile, err := ioutil.TempFile(tempDir, tempFileSchedulerName)
	require.NoErrorf(t, err, "unable to create temp file")
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			require.NoErrorf(t, err, "unable to delete temp file %s", tempFile.Name())
		}
	}()

	if _, err = tempFile.WriteString(testConfigSchedulerContent); err != nil {
		require.NoError(t, err, "unable to write content to file %s", tempFile.Name())
	}

	if err := tempFile.Close(); err != nil {
		require.NoErrorf(t, err, "unable to close temp file %s", tempFile.Name())
	}

	config := NewConfigScheduler()
	err = config.BuildConfigScheduler(tempFile.Name())
	require.NoError(t, err)

	require.Equal(t, "INFO1", config.Logger.Level)
	require.Equal(t, "log_schedule1.log", config.Logger.Path)
	require.Equal(t, "127.0.0.1", config.Server.Host)
	require.Equal(t, "8080", config.Server.HttpPort)
	require.Equal(t, "50051", config.Server.GrpcPort)
	require.Equal(t, "SQL", config.Storage)
	require.Equal(t, "5432", config.Database.Port)
	require.Equal(t, "127.0.0.1", config.Database.Host)
	require.Equal(t, "postgres1", config.Database.Username)
	require.Equal(t, "password1", config.Database.Password)
	require.Equal(t, "calendar", config.Database.Name)
}
