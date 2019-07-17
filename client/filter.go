package client

import (
	"strings"

	"github.com/rburmorrison/hoist/types"
)

// Filter returns a list of repository summaries from
// a registry that contain a string.
func Filter(text string) ([]types.RepositorySummary, error) {
	summaries, err := Repos()
	if err != nil {
		return nil, err
	}

	var filteredSummaries []types.RepositorySummary
	for _, summary := range summaries {
		if strings.Contains(summary.Name, text) {
			filteredSummaries = append(filteredSummaries, summary)
		}
	}

	return filteredSummaries, nil
}
