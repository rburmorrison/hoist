package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rburmorrison/hoist/types"
)

func configSetAll(config types.Configuration) error {
	bs, err := json.Marshal(config)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(types.ConfigFilePath, append(bs, '\n'), 770); err != nil {
		return err
	}

	return nil
}

func configSetKey(key string, value interface{}) error {
	config, err := configFetchAll()
	if err != nil {
		return err
	}

	config[key] = value
	return configSetAll(config)
}

// ConfigSetAddress will set the address setting for
// hoist. An error may occur if the configuration
// file is unavailable.
func ConfigSetAddress(address string) error {
	return configSetKey("address", address)
}

// ConfigSetMode will set the mode setting for hoist.
// An error may occur if the configuration file is
// unavailable.
func ConfigSetMode(mode types.Mode) error {
	return configSetKey("mode", mode)
}
