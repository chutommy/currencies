package config

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

// Config defines the service settings.
type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// GetConfig reads the config file and returns its configuration.
func GetConfig(fileName string) (*Config, error) {

	// from file
	bs, err := ioutil.ReadFile(filepath.Join(rootDir(), fileName))
	if err != nil {
		return nil, fmt.Errorf("could not read config file: %w", err)
	}

	// decode config file
	var cfg Config
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return nil, fmt.Errorf("could not decode config file content: %w", err)
	}

	return &cfg, nil
}

// rootDir returns the path of the root directory.
func rootDir() string {
	_, caller, _, _ := runtime.Caller(0)
	dir := path.Dir(caller)
	return filepath.Dir(dir)
}
