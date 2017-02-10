package jira

import (
	"encoding/json"
	"io/ioutil"

	homedir "github.com/mitchellh/go-homedir"
)

// Config struct for storing JSON data file
type Config struct {
	Host string
}

// GetConfig reads the config file from the possible locations and then
// returns a pointer to a new config object.
func GetConfig() (*Config, error) {
	path, err := homedir.Expand("~/.jira/config.json")
	if err != nil {
		return nil, err
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	json.Unmarshal(file, &config)
	return &config, nil
}

func (c Config) String() string {
	return c.Host
}
