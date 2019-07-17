// Package client defines the client functionality
// for hoist.
package client

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/rburmorrison/hoist/types"
)

var (
	// ErrConfigNotFound occurs when trying to access
	// a configuration option that is not defined.
	ErrConfigNotFound = errors.New("config key not found")

	// ErrImageNotFound occurs when the search function
	// is unable to find the provided image.
	ErrImageNotFound = errors.New("image not found")
)

// sendRequest sends a GET request to the registry
// and returns the body of the response.
func sendRequest(route string) ([]byte, error) {
	// Fetch configuraiton
	mode, err := ConfigFetchMode()
	if err != nil {
		return nil, err
	}

	address, err := ConfigFetchAddress()
	if err != nil {
		return nil, err
	}

	// Generate a header from the mode
	var header string
	switch mode {
	case types.ModeHTTP:
		header = "http://"
	case types.ModeHTTPS:
		header = "https://"
	}

	// Make a request to the API
	path := header + address + route
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response
	return ioutil.ReadAll(res.Body)
}
