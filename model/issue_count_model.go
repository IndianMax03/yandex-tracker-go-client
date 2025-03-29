// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// IssueCountRequest describes request to get the number of issues
type IssueCountRequest struct {
	// Issue filtering parameters.
	// In the parameter, you can specify the name of any field and the value by which filtering will be performed.
	Filter map[string]any `json:"filter"`
	// Filter in query language.
	Query string `json:"query"`
}

// IssueCountResponse describes response countains number of issues
type IssueCountResponse int
