package core

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const (
	TempConfigFile = "temp-config.yaml"
)

func TestBazookaConfig_FromString (t *testing.T) {
	var err error = nil
	in := `
database:
  dialect: sqlite
  user: fake-db-user
`

	c := BazookaConfig{}
	err = c.FromString(in)
	assert.NoError(t, err)

	assert.Equal(t, "sqlite", c.DbArgs.Dialect)
	assert.Equal(t, "fake-db-user", c.DbArgs.User)
}


func TestBazookaConfig_FromFile(t *testing.T) {
	var err error = nil
	in := `
database:
  dialect: sqlite
  user: fake-db-user
`
	// Write temp config yaml.
	err = ioutil.WriteFile(TempConfigFile, []byte(in), 0644)
	assert.NoError(t, err)

	c := BazookaConfig{}
	err = c.FromFile(TempConfigFile)
	assert.NoError(t, err)

	assert.Equal(t, "sqlite", c.DbArgs.Dialect)
	assert.Equal(t, "fake-db-user", c.DbArgs.User)

	// Cleanup
	_ = os.Remove(TempConfigFile)
}
