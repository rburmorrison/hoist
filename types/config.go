package types

// Mode represents the way that requests are sent to
// a Docker registry.
type Mode int

const (
	// ModeHTTP sends requests to the registry using HTTP
	ModeHTTP Mode = iota

	// ModeHTTPS sends requests to the registry using
	// HTTPS
	ModeHTTPS
)

func (m Mode) String() string {
	switch m {
	case ModeHTTP:
		return "HTTP"
	case ModeHTTPS:
		return "HTTPS"
	}
	return ""
}

// Configuration represents the settings for hoist.
type Configuration map[string]interface{}
