package client

import (
	"encoding/json"

	"github.com/rburmorrison/hoist/types"
)

// Tags returns the tags for a repository in a
// registry.
func Tags(repository string) ([]string, error) {
	bs, err := sendRequest("/v2/" + repository + "/tags/list")
	if err != nil {
		return nil, err
	}

	var tres types.TagsResponse
	if err = json.Unmarshal(bs, &tres); err != nil {
		return nil, err
	}

	return tres.Tags, nil
}
