// The model package contains an entities for exchanging information with the Yandex Tracker API
package model

type PageRequest struct {
	PerPage int
	Page    int
}

type PageResponse struct {
	TotalPages int
	TotalCount int
}
