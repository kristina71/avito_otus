package config

import (
	"os"

	"gopkg.in/yaml.v2" //nolint:typecheck
)

type RMQConfig struct {
	Logger struct {
		Level string `yaml:"level"`
		Path  string `yaml:"path"`
	} `yaml:"logger"`
	RMQ struct {
		Uri   string `yaml:"uri"`
		Queue string `yaml:"queue"`
	} `yaml:"rmq"`
}

func NewConfigRMQ() RMQConfig {
	return RMQConfig{}
}

func (config *RMQConfig) BuildConfigRMQ(path string) error {
	// Open the configuration file
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	d := yaml.NewDecoder(f) // nolint:typecheck

	if err = d.Decode(&config); err != nil {
		return err
	}

	return nil
}
