// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// Queue describes the queue the issue is in (request)
type Queue struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

// Board has no info: https://yandex.ru/support/tracker/ru/concepts/issues/search-issues
type Board struct {
	ID   any    `json:"id"`
	Name string `json:"name"`
}

// ObjectBaseRequest describes the basic structure of a request object
type ObjectBaseRequest struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

// ObjectBaseResponse describes the basic structure of a request object
type ObjectBaseResponse struct {
	Self    string `json:"self"`
	ID      string `json:"id"`
	Display string `json:"display"`
}

// IssueParent describes parent of issue
type IssueParent struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

// UpdatedBy describes the user who updated the entity last
type UpdatedBy ObjectBaseResponse

// IssueSprint describes issue sprint
type IssueSprint []ObjectBaseResponse

// IssueType describes issue type
type IssueType struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

// IssuePriority describes issue priority
type IssuePriority struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

// IssueFollowers describes issue followers
type IssueFollowers []ObjectBaseResponse

// CreatedBy describes the user who created the entity
type CreatedBy ObjectBaseResponse

// IssueAssignee describes user who is the issue performer
type IssueAssignee ObjectBaseResponse

// IssueQueue describes the queue the issue is in (response)
type IssueQueue struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

// IssueStatus describes the issue status
type IssueStatus struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

// IssueStatusType describes the issue status type
type IssueStatusType struct {
	ID      string `json:"id"`
	Display string `json:"display"`
	Key     string `json:"key"`
}

// IssuePreviousStatus describes previous issue status
type IssuePreviousStatus struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

// Attachment describes issue attachment
type Attachment ObjectBaseResponse

// Transition describes possible issue transition
type Transition ObjectBaseResponse
