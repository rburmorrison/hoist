// Package client defines the client functionality
// for hoist.
package client

import "errors"

var (
	// ErrConfigNotFound occurs when trying to access
	// a configuration option that is not defined.
	ErrConfigNotFound = errors.New("config key not found")
)
