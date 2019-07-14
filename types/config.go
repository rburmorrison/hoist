package types

type Mode int

var (
	// ModeHTTP sends requests to the registry using HTTP
	ModeHTTP Mode = iota

	// ModeHTTPS sends requests to the registry using
	// HTTPS
	ModeHTTPS Mode
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
