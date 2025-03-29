// The model package contains an entities for exchanging information with the Yandex Tracker API
package model

type Queue struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

type Board struct {
	ID   any    `json:"id"`
	Name string `json:"name"`
}

type ObjectBaseRequest struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

type ObjectBaseResponse struct {
	Self    string `json:"self"`
	ID      string `json:"id"`
	Display string `json:"display"`
}

type IssueParent struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

type IssueUpdatedBy ObjectBaseResponse

type IssueSprint []ObjectBaseResponse

type IssueType struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

type IssuePriority struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

type IssueFollowers []ObjectBaseResponse

type IssueCreatedBy ObjectBaseResponse

type IssueAssignee ObjectBaseResponse

type IssueQueue struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

type IssueStatus struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

type IssueStatusType struct {
	ID      string `json:"id"`
	Display string `json:"display"`
	Key     string `json:"key"`
}

type IssuePreviousStatus struct {
	ObjectBaseResponse
	Key string `json:"key"`
}
