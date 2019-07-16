package types

// CatalogResponse represents the response from the
// /v2/_catalog route.
type CatalogResponse struct {
	Repositories []string `json:"repositories"`
}

// TagsResponse represents the response from the
// /v2/<name>/tags/list route.
type TagsResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}
