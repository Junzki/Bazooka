package core

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	DatabaseUrlFormat = "%s://"
)

type DatabaseArgs struct {
	Dialect  string `yaml:"dialect"`
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Charset  string `yaml:"charset"`
}


type BazookaConfig struct {
	DbArgs	DatabaseArgs	`yaml:"database"`
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
