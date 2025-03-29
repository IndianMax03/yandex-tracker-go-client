// The model package contains an entities for exchanging information with the Yandex Tracker API
package model

// Request to get the number of issues
type IssueCountRequest struct {
	// Issue filtering parameters.
	// In the parameter, you can specify the name of any field and the value by which filtering will be performed.
	Filter map[string]any `json:"filter"`
	// Filter in query language.
	Query string `json:"query"`
}

// Number of issues
type IssueCountResponse int
