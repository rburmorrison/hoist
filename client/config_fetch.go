package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rburmorrison/hoist/types"
)

func configFetchAll() (types.Configuration, error) {
	bs, err := ioutil.ReadFile(types.ConfigFilePath)
	if err != nil {
		return nil, err
	}

	var config types.Configuration
	if err = json.Unmarshal(bs, &config); err != nil {
		return nil, err
	}

	return config, nil
}

func configFetchKey(key string) (interface{}, error) {
	config, err := configFetchAll()
	if err != nil {
		return nil, err
	}

	val, ok := config[key]
	if !ok {
		return nil, ErrConfigNotFound
	}

	return val, nil
}

// ConfigFetchAddress will set the address setting for
// hoist. An error may occur if the configuration
// file is unavailable.
func ConfigFetchAddress() (string, error) {
	val, err := configFetchKey("address")
	if err != nil {
		return "", err
	}

	return val.(string), nil
}

// ConfigFetchMode will set the mode setting for
// hoist. An error may occur if the configuration
// file is unavailable.
func ConfigFetchMode() (types.Mode, error) {
	val, err := configFetchKey("mode")
	if err != nil {
		return -1, err
	}

	return val.(types.Mode), nil
}
