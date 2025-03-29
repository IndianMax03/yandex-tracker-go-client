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

type ObjectBase struct {
	Self    string `json:"self"`
	ID      string `json:"id"`
	Display string `json:"display"`
}

type IssueParent struct {
	ObjectBase
	Key string `json:"key"`
}

type IssueUpdatedBy ObjectBase

type IssueSprint []ObjectBase

type IssueType struct {
	ObjectBase
	Key string `json:"key"`
}

type IssuePriority struct {
	ObjectBase
	Key string `json:"key"`
}

type IssueFollowers []ObjectBase

type IssueCreatedBy ObjectBase

type IssueAssignee ObjectBase

type IssueQueue struct {
	ObjectBase
	Key string `json:"key"`
}

type IssueStatus struct {
	ObjectBase
	Key string `json:"key"`
}

type IssueStatusType struct {
	ID      string `json:"id"`
	Display string `json:"display"`
	Key     string `json:"key"`
}

type IssuePreviousStatus struct {
	ObjectBase
	Key string `json:"key"`
}
