package config

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

var SECTIONDECLIMNOTMATCH = errors.New("SECTION_DECLIM_NOT_MATCH")

type configSection struct {
	sectionName string
	confItems   map[string]interface{}
}

// Config :
type Config struct {
	section map[string]*configSection
}

// NewConfig : instantiate Config according to the configure file(full path).
func NewConfig(fullPath string) (*Config, error) {
	config := new(Config)

	if err := config.ReloadFromFile(fullPath); err != nil {
		return nil, err
	}

	return config, nil

}

// ReloadFromFile : reload conf data from a config file
func (config *Config) ReloadFromFile(fullPath string) error {
	fp, err := os.Open(fullPath)

	if err != nil {
		return err
	}

	return config.ReloadFromReader(fp)
}

// ReloadFromReader : reload conf data from an io.reader
func (config *Config) ReloadFromReader(input io.Reader) error {
	var currSectName string // current section name

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {

	}
}

// comment line
func isCommentLine(line string, delim string) bool {
	if strings.HasPrefix(line, delim) {
		return true
	}

	return false
}

// empty line
func isEmptyLine(line string) bool {
	if len(line) == 0 {
		return true
	}

	return false
}

// section line, below is an example
// [ db ], then caller is
// sectName, ok, err:=isSectionLine(line, "[", "]")
func isSectionLine(line string, left string, right string) (string, bool, error) {
	if strings.HasPrefix(line, left) {
		if !strings.HasSuffix(line, right) {
			return "", false, errors.New("")
		}
	}
}
