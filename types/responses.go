package types

// CatalogResponse represents the response from the
// /v2/_catalog route.
type CatalogResponse struct {
	Repositories []string `json:"repositories"`
}
