package types

// RepositorySummary represents a description of a
// registry in a Docker registry.
type RepositorySummary struct {
	Name     string
	TagCount uint
}
