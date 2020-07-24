package config

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetConfig(t *testing.T) {

	tests := []struct {
		name     string
		fileName string
		err      string
	}{
		{
			name:     "ok",
			fileName: "config/tests/config_0.yaml",
			err:      "",
		},
		{
			name:     "invalid config content",
			fileName: "config/tests/config_1.yaml",
			err:      "could not decode config file content",
		},
		{
			name:     "non existing file",
			fileName: "config/tests/config_2.yaml",
			err:      "could not read config file",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t1 *testing.T) {

			cfg, err := GetConfig(test.fileName)
			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.err)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				cfg.Host = "127.0.0.1"
				cfg.Port = 10502
				assert.Equal(t1, "", test.err)
			}
		})
	}
}
