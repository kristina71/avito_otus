package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
/*type Config struct {
	Logger       LoggerConf   `yaml:"logger"`
	Server       ServerConf   `yaml:"server"`
	DatabaseConf DatabaseConf `yaml:"database"`
	Storage      string       `yaml:"storage"`
}

*/

type Config struct {
	Logger struct {
		Level string
		File  string `yaml:"file"`
	} `yaml:"logger"`
	Server struct {
		Host     string `yaml:"host"`
		HTTPPort string `yaml:"httpPort"`
		GrpcPort string `yaml:"grpcPort"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"sslMode"`
	} `yaml:"database"`
	Storage string `yaml:"storage"`
}

func NewConfig(configFile string) (Config, error) {
	config := Config{}

	v := viper.New()

	if configFile != "" {
		fmt.Println(configFile)
		v.SetConfigFile(configFile)
		err := v.ReadInConfig()
		if err != nil {
			return config, fmt.Errorf("failed to read configuration: %w", err)
		}
	} else {
		configure(v)
	}

	fmt.Println(config)
	if err := v.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	fmt.Println(config)
	if err := config.Validate(); err != nil {
		return config, fmt.Errorf("failed to validate configuration: %w", err)
	}

	return config, nil
}

func configure(v *viper.Viper) {
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	v.SetDefault("logger.level", "INFO")
	v.SetDefault("logger.file", "log.log")

	v.SetDefault("server.host", "127.0.0.1")
	v.SetDefault("server.httpPort", "8080")
	v.SetDefault("server.grpcPort", "50051")

	v.SetDefault("database.host", "127.0.0.1")
	v.SetDefault("database.port", "5432")
	v.SetDefault("database.username", "postgres")
	v.SetDefault("database.password", "password")
	v.SetDefault("database.name", "calendar")
	v.SetDefault("database.SSLMode", "disable")

	v.SetDefault("storage", "SQL")
}

func (c Config) Validate() error {
	fmt.Println(c.Server)
	if c.Server.Host == "" {
		return errors.New("http app server host is required")
	}

	if c.Server.HTTPPort == "" {
		return errors.New("http app server port is required")
	}

	if c.Server.GrpcPort == "" {
		return errors.New("internalgrpc app server port is required")
	}

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
