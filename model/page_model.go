// Model package contains an entities for exchanging information with the Yandex Tracker API
package model

// PageRequest describes pagination request headers
type PageRequest struct {
	PerPage int
	Page    int
}

// PageResponse describes pagination response headers
type PageResponse struct {
	TotalPages int
	TotalCount int
}
