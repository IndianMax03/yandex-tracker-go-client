// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// PageRequest describes pagination request headers
type PageRequest struct {
	// Objects per page
	PerPage int
	// Page number
	Page int
	// ID parameter of the object after which the requested page will begin
	FromID int
}

// PageResponse describes pagination response headers
type PageResponse struct {
	// Total number of pages
	TotalPages int
	// Total number of objects
	TotalCount int
	// ID parameter of the object that was the last one on the requested page
	LastID int
}
