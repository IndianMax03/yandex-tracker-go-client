// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// PriorityResponse describes an object that contains priority information
type PriorityResponse struct {
	// The address of the API resource that contains priority information.
	Self string `json:"self"`
	// Priority identifier.
	ID int `json:"id"`
	// Priority key.
	Key string `json:"key"`
	// Priority version.
	Version int `json:"version"`
	// Display name of the priority.
	// If localized=false is passed in the request, this parameter will contain duplicates of the name in other languages.
	// localized=true -> string
	// localized=false -> map[string]string
	Name any `json:"name"`
	// No info: https://yandex.ru/support/tracker/ru/concepts/issues/get-priorities
	Description string `json:"description"`
	// Priority weight. The parameter affects the order in which the priority is displayed in the interface.
	Order int `json:"order"`
}
