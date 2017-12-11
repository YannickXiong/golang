package config

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// error code
var (
	ErrSectionDelimNotMatch = errors.New("SECTION_DELIM_NOT_MATCH")
	ErrKeyValueFormat       = errors.New("KEY_VALUE_FORMAT_ERROR")
)

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
	config.section = make(map[string]*configSection)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		// trim comment info at the end of a stament
		line = trimComment(line, "#")
		line = trimComment(line, ";")

		// comment line, support # && ;
		if isCommentLine(line, "#") || isCommentLine(line, ";") {
			continue
		}

		// empty line
		if isEmptyLine(line) {
			continue
		}

		sectName, hasSection, err := isSectionLine(line, "[", "]")
		if hasSection {
			currSectName = sectName
			confItems := make(map[string]interface{})

			// section not exist
			if _, ok := config.section[currSectName]; !ok {
				config.section[currSectName] = new(configSection)
				config.section[currSectName].confItems = confItems
				config.section[currSectName].sectionName = currSectName
			}
		} else if err != nil {
			// fmt.Printf("errCode = %s, errMsg = %s", err, "section name must be surrounding by '[' and ']'.")
			return err
		}

		key, value, hasKeyValue, err := isKeyValueLine(line, "=")
		if hasKeyValue {

			// a key-value statment dosen't belong to any section, it will belong to "global" section.
			if currSectName == "" {
				currSectName = "global"
				if _, ok := config.section[currSectName]; !ok {
					confItems := make(map[string]interface{})
					config.section[currSectName] = new(configSection)
					config.section[currSectName].confItems = confItems
					config.section[currSectName].sectionName = currSectName
				}
			}

			// key value doesn't exist.
			if _, ok := config.section[currSectName].confItems[key]; !ok {
				config.section[currSectName].confItems[key] = value
			}
		} else if err != nil {
			// ftm.Printf("errCode = %s, errMsg = %s", err, "key and value must be connected by '='.")
			return err
		}

	}

	return nil
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

// section line, below is an example, [ db ], then caller is
// sectName, ok, err:=isSectionLine(line, "[", "]")
// returns:
//		"", false, error : when sectionName is not surrounding with left & right delim, eg. [ db
//	    sectName, true, nil : when line is a section line, eg. [ db ]
// 	    "", false, nil : when line is not a section line, may be a key-value line, eg. dbPort = 3306
func isSectionLine(line string, left string, right string) (string, bool, error) {
	var sectName string
	var isSection bool

	if strings.HasPrefix(line, left) {
		if !strings.HasSuffix(line, right) {
			// return "", false, ErrSectionDelimNotMatch
			return sectName, isSection, ErrSectionDelimNotMatch
		}

		sectName := line[1 : len(line)-1]
		sectName = strings.TrimSpace(sectName) // [ db ] => db, not " db "
		isSection = true

		// return sectName, true, nil
		return sectName, isSection, nil
	}

	// return sectName, false, nil
	return sectName, isSection, nil
}

// key-value line, below is an example, dbPort = 3306, then caller is
// key, value, ok, err:=isKeyValueLine(line, "=")
// returns:
//		"","", false, error : when a key has no values or values has no key, eg. dbPort = or = 3306
//	    key, value, true, nil : when line is an key-value line, eg. dbPort = 3306
// 	    "", "", false, nil : when line is not a key-value line, may be a section line, eg. [ db ]
func isKeyValueLine(line string, delim string) (string, string, bool, error) {
	var key string
	var value string
	var isKeyValue bool

	if !strings.Contains(line, delim) {
		return key, value, isKeyValue, nil
	}

	equalIndex := strings.Index(line, delim)

	if equalIndex <= 0 {
		return key, value, isKeyValue, ErrKeyValueFormat
	}

	key = strings.ToLower(strings.TrimSpace(line[:equalIndex]))
	value = strings.ToLower(strings.TrimSpace(line[equalIndex+1:]))
	isKeyValue = true

	return key, value, isKeyValue, nil
}

// trimComment : comment at the end of a statment, eg. dbPort = 3306  # db port
func trimComment(line string, delim string) string {

	commentIndex := strings.Index(line, delim)
	if commentIndex > 0 {
		// remove the comment info
		// if not tripSpace, [ aaa ] # will return " [ aaa ] ",
		// then isSectionLine will fail with strings.HasSuffix(line, right) failed.
		return strings.TrimSpace(line[:commentIndex])
	}

	return line
}

// Configer : interface of Config
type Configer interface {
	ToString() string
	GetSections() []string
	GetIntegerFromSection(sectionName string, key string, defaultValue int64) int64
	GetStringrFromSection(sectionName string, key string, defaultValue string) string
	GetBoolFromSection(sectionName string, key string, defaultValue bool) bool
	GetFloatFromSection(sectionName string, key string, defaultValue float64) float64
}

// ToString : convert Config object to string
func (config *Config) ToString() string {
	var output string
	for _, section := range config.GetSections() {
		output += fmt.Sprintf("section = %s\n", section)
		for key, value := range config.section[section].confItems {
			output += fmt.Sprintf("    %s = %s\n", key, value)
		}
	}

	return output
}

// GetSections : get all section names
func (config *Config) GetSections() []string {

	sections := make([]string, 0)

	for key := range config.section {
		sections = append(sections, key)
	}

	return sections
}
