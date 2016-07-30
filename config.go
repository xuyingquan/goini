/**
 * Read the configuration file
 *
 * @copyright    (C) 2016  xuyingquan
 * @lastmodify   2016-7-29
 * @website		http://www.statacloud.com
 *
 */

package goini

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	info map[string]map[string]string // configuration information
}

func NewConfig(path string) *Config {
	c := &Config{info: make(map[string]map[string]string)}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()

	var section string
	buf := bufio.NewReader(file)

	for {
		l, err := buf.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				CheckErr(err)
			} else {
				break
			}
		}
		line := strings.TrimSpace(l)

		switch {
		case len(line) == 0 || line[0] == '#':
			continue
		case line[0] == '[' && line[len(line)-1] == ']':
			section = strings.TrimSpace(line[1 : len(line)-1])
			c.info[section] = make(map[string]string)
		default:
			i := strings.IndexAny(line, "=")
			key := strings.TrimSpace(line[:i])
			value := strings.TrimSpace(line[i+1 : len(line)])
			c.info[section][key] = value
		}
	}

	return c
}

func (c *Config) GetString(section string, field string) (string, error) {
	value, exists := c.info[section][field]
	if exists {
		return value, nil
	}
	err := errors.New("the field is not exist.")
	return "", err
}

func (c *Config) GetInt(section string, field string) (int, error) {
	value, exists := c.info[section][field]
	if exists {
		return strconv.Atoi(value)
	}
	err := errors.New("the field is not exist.")
	return -1, err
}

func (c *Config) GetBool(section string, field string) (bool, error) {
	value, exists := c.info[section][field]
	if exists {
		return strconv.ParseBool(value)
	}
	err := errors.New("the field is not exist.")
	return false, err
}

func (c *Config) GetFloat(section string, field string) (float64, error) {
	value, exists := c.info[section][field]
	if exists {
		return strconv.ParseFloat(value, 64)
	}
	err := errors.New("the field is not exist.")
	return 0, err
}

func CheckErr(err error) string {
	if err != nil {
		return fmt.Sprintf("Error is :'%s'", err.Error())
	}
	return "Notfound this error"
}
