package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rburmorrison/hoist/types"
)

// Reps returns a list of repository summaries from
// a registry.
func Reps() ([]types.RepositorySummary, error) {
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
	path := header + address + "/v2/_catalog"
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode the response
	bs, err := ioutil.ReadAll(res.Body)
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

		summaries = append(summaries, summary)
	}

	return summaries, nil
}
