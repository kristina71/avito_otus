package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	tempDirectoryRmqPath = "./"
	tempDirRmqName       = "tempdir"
	tempFileRmqName      = "tempconfig*.yaml"
	testConfigRmqContent = `
logger:
  level: "INFO"
  path: "log_rmq.log"
rmq:
  uri: "amqp://guest:guest@localhost:5672/"
  queue: "calendar"`
)

func TestConfigRmgReading(t *testing.T) {
	// create temp directory
	tempDir, err := ioutil.TempDir(tempDirectoryRmqPath, tempDirRmqName)
	require.NoErrorf(t, err, "unable to create temp directory")
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			require.NoErrorf(t, err, "unable to delete temp dir %s", tempDir)
		}
	}()
	// crete temp file in temp directory
	tempFile, err := ioutil.TempFile(tempDir, tempFileRmqName)
	require.NoErrorf(t, err, "unable to create temp file")
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			require.NoErrorf(t, err, "unable to delete temp file %s", tempFile.Name())
		}
	}()

	if _, err = tempFile.WriteString(testConfigRmqContent); err != nil {
		require.NoError(t, err, "unable to write content to file %s", tempFile.Name())
	}

	if err := tempFile.Close(); err != nil {
		require.NoErrorf(t, err, "unable to close temp file %s", tempFile.Name())
	}

	config := NewConfigRMQ()
	err = config.BuildConfigRMQ(tempFile.Name())
	require.NoError(t, err)

	require.Equal(t, "INFO", config.Logger.Level)
	require.Equal(t, "log_rmq.log", config.Logger.Path)
	require.Equal(t, "amqp://guest:guest@localhost:5672/", config.RMQ.Uri)
	require.Equal(t, "calendar", config.RMQ.Queue)
}
