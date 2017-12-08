package config

import (
	"io"
	"os"
)

type confItemNameType string

type confItemValueType interface{}

type confSectNameType string

type configSection struct {
	sectionName confSectNameType
	confItems   map[confItemNameType]confItemValueType
}

type Config struct {
	section map[confSectNameType]*configSection
}

func NewConfig(fullPath string) (*Config, error) {
	config := new(Config)
	fp, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}

	defer fp.Close()

	return config.Reload(f)
}

func (config *Config) Reload(input io.Reader) error {

}
