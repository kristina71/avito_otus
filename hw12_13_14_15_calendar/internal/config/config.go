package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger       LoggerConf `yaml:"logger"`
	Server       ServerConf `yaml:"server"`
	DatabaseConf DatabaseConf
}

/*
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"SSLMode"`
	} `yaml:"database"`
	Storage string `yaml:"storage"`
}*/

type LoggerConf struct {
	Level string
	File  string
}

type ServerConf struct {
	Host     string `yaml:"host"`
	HTTPPort string `yaml:"http_port"`
	GrpcPort string `yaml:"grpc_port"`
}

type DatabaseConf struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"SSLMode"`
	} `yaml:"database"`
	Storage string `yaml:"storage"`
}

func NewConfig(configFile string) (Config, error) {
	config := Config{}

	v := viper.New()

	configure(v)

	if configFile != "" {
		v.SetConfigFile(configFile)
		err := v.ReadInConfig()
		if err != nil {
			return config, fmt.Errorf("failed to read configuration: %w", err)
		}
	}

	if err := v.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	if err := config.Validate(); err != nil {
		return config, fmt.Errorf("failed to validate configuration: %w", err)
	}

	return config, nil
}

func configure(v *viper.Viper) {
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	v.SetDefault("logger.level", "INFO")

	v.SetDefault("server.host", "127.0.0.1")
	v.SetDefault("server.httpPort", "8080")
	v.SetDefault("server.grpcPort", "8081")

	v.SetDefault("database.inmem", true)
}

func (c ServerConf) Validate() error {
	if c.Host == "" {
		return errors.New("http app server host is required")
	}

	if c.HTTPPort == "" {
		return errors.New("http app server port is required")
	}

	if c.GrpcPort == "" {
		return errors.New("grpc app server port is required")
	}

	return nil
}

func (c Config) Validate() error {
	if err := c.Server.Validate(); err != nil {
		return err
	}

	if err := c.Database.Validate(); err != nil {
		return err
	}

	return nil
}

func (c DatabaseConf) Validate() error {
	if c.Database.Password == "" {
		return errors.New("db password is required")
	}

	if c.Database.SSLMode == "" {
		return errors.New("SSLMode is required")
	}

	if c.Database.Host == "" {
		return errors.New("db host is required")
	}

	if c.Database.Name == "" {
		return errors.New("db name is required")
	}

	if c.Database.Username == "" {
		return errors.New("db username is required")
	}

	if c.Database.Port == "" {
		return errors.New("db port is required")
	}

	return nil
}
