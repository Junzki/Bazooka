package core

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type DatabaseConf struct {
	Dialect  string            `yaml:"dialect"`
	Host     string            `yaml:"host"`
	Port     uint              `yaml:"port"`
	User     string            `yaml:"user"`
	Password string            `yaml:"password"`
	Name     string            `yaml:"name"`
	Charset  string            `yaml:"charset"`
	Extra    map[string]string `yaml:"extra"`
}

func (c DatabaseConf) GetConnString() (string, string, error) {
	dialect := strings.ToLower(c.Dialect)

	switch dialect {
	case "postgres":
		conn, err := c.getPostgresConn()
		return dialect, conn, err
	default:
		err := errors.New(fmt.Sprintf("dialect %s not supported", dialect))
		return dialect, "", err
	}
}

func (c DatabaseConf) getPostgresConn() (string, error) {
	if "" == c.Name {
		return "", errors.New("dbname not specified")
	}

	var conn = ""
	if "" != c.Host {
		conn += fmt.Sprintf("host=%s ", c.Host)
	}

	if 0 != c.Port {
		conn += fmt.Sprintf("port=%d ", c.Port)
	}

	if "" != c.User {
		conn += fmt.Sprintf("user=%s password=%s ", c.User, c.Password)
	}

	conn += fmt.Sprintf("dbname=%s ", c.Name)

	// Apply extra configurations.
	for k, v := range c.Extra {
		conn += fmt.Sprintf("%s=%s ", k, v)
	}

	return conn, nil
}

type BazookaConfig struct {
	Debug    bool         `yaml:"debug"`
	Listen   string       `yaml:"listen"`
	Port     uint         `yaml:"port"`
	Database DatabaseConf `yaml:"database"`
}

func (c *BazookaConfig) FromFile(p string) error {
	// Check if file exists.
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return err
	}

	f, err := os.Open(p)
	if nil != err {
		return err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if nil != err {
		return err
	}

	err = c.FromString(buf)
	if nil != err {
		return err
	}

	return nil
}

func (c *BazookaConfig) FromString(i interface{}) error {
	var in []byte
	switch v := i.(type) {
	case string:
		in = []byte(v)
	case []byte:
		in = v
	default:
		return errors.New("bad input format")
	}

	err := yaml.Unmarshal(in, c)
	return err
}

var config *BazookaConfig

func GetConfig() *BazookaConfig {
	if nil != config {
		return config
	}

	config = &BazookaConfig{
		Listen: "127.0.0.1",
		Port:   8081,
		Debug:  true,
	}
	return config
}
