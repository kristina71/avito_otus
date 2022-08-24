package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	tempDirectoryPath = "./"
	tempDirName       = "tempdir"
	tempFileName      = "tempconfig*.yaml"
	testConfigContent = `
logger:
  level: "INFO1"
  file: "log1.log"
server:
  httpPort: "8081"
  grpcPort: "50011"
  host: "121.0.0.1"
database:
  host: "121.0.0.1"
  port: "5431"
  username: "postgres1"
  password: "password1"
  name: "calendar1"
  SSLMode: "disable1"
storage: "SQL1"`
)

func TestConfigDefault(t *testing.T) {
	config, err := NewConfig("")
	require.NoError(t, err)

	require.Equal(t, "INFO", config.Logger.Level)
	require.Equal(t, "log.log", config.Logger.File)
	require.Equal(t, "8080", config.Server.HTTPPort)
	require.Equal(t, "50051", config.Server.GrpcPort)
	require.Equal(t, "SQL", config.Storage)
	require.Equal(t, "5432", config.Database.Port)
	require.Equal(t, "127.0.0.1", config.Database.Host)
	require.Equal(t, "postgres", config.Database.Username)
	require.Equal(t, "password", config.Database.Password)
	require.Equal(t, "calendar", config.Database.Name)
}

func TestConfigReading(t *testing.T) {
	// create temp directory
	tempDir, err := ioutil.TempDir(tempDirectoryPath, tempDirName)
	require.NoErrorf(t, err, "unable to create temp directory")
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			require.NoErrorf(t, err, "unable to delete temp dir %s", tempDir)
		}
	}()
	// crete temp file in temp directory
	tempFile, err := ioutil.TempFile(tempDir, tempFileName)
	require.NoErrorf(t, err, "unable to create temp file")
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			require.NoErrorf(t, err, "unable to delete temp file %s", tempFile.Name())
		}
	}()

	if _, err = tempFile.WriteString(testConfigContent); err != nil {
		require.NoError(t, err, "unable to write content to file %s", tempFile.Name())
	}

	if err := tempFile.Close(); err != nil {
		require.NoErrorf(t, err, "unable to close temp file %s", tempFile.Name())
	}

	fmt.Println(tempFile.Name())
	config, err := NewConfig(tempFile.Name())
	require.NoError(t, err)

	require.Equal(t, "INFO1", config.Logger.Level)
	require.Equal(t, "log1.log", config.Logger.File)
	require.Equal(t, "8081", config.Server.HTTPPort)
	require.Equal(t, "50011", config.Server.GrpcPort)
	require.Equal(t, "SQL1", config.Storage)
	require.Equal(t, "5431", config.Database.Port)
	require.Equal(t, "121.0.0.1", config.Database.Host)
	require.Equal(t, "postgres1", config.Database.Username)
	require.Equal(t, "password1", config.Database.Password)
	require.Equal(t, "calendar1", config.Database.Name)
}
