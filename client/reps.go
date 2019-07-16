package client

import (
	"encoding/json"

	"github.com/rburmorrison/hoist/types"
)

// Reps returns a list of repository summaries from
// a registry.
func Reps() ([]types.RepositorySummary, error) {
	bs, err := sendRequest("/v2/_catalog")
	if err != nil {
		return nil, err
	}

	var cres types.CatalogResponse
	if err = json.Unmarshal(bs, &cres); err != nil {
		return nil, err
	}

	// Summarize repositories
	var summaries []types.RepositorySummary
	for _, repository := range cres.Repositories {
		var summary types.RepositorySummary
		summary.Name = repository

		// Count tags
		tags, err := Tags(summary.Name)
		if err != nil {
			return nil, err
		}

		summary.TagCount = uint(len(tags))
		summaries = append(summaries, summary)
	}

	return summaries, nil
}
