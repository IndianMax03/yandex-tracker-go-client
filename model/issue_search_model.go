// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// IssueSearchRequest describes request to get filtered and sorted issues. The request does not imply a combination of parameters.
// According to priority:
// 1. queue;
// 2. keys;
// 3. filter + order;
// 4. query.
type IssueSearchRequest struct {
	// Queue.
	Queue string `json:"queue,omitempty"`
	// List of task keys.
	Keys []string `json:"keys,omitempty"`
	// Issue filtering parameters.
	// In the parameter, you can specify the name of any field and the value by which filtering will be performed.
	Filter map[string]any `json:"filter,omitempty"`
	// The direction and field of issue sorting (only in combination with the Filter parameter).
	// The value is specified in the format [+/-]<field_key>. The + or - sign indicates the sorting direction.
	Order string `json:"order,omitempty"`
	// Filter in query language.
	Query string `json:"query,omitempty"`
}

// IssueResponse describes response contains array of objects that containing information about issues
type IssueResponse struct {
	// The address of the API resource that contains information about the issue.
	Self string `json:"self"`
	// Issue ID.
	ID string `json:"id"`
	// Issue Key.
	Key string `json:"key"`
	// Issue version. Each change to the issue parameters increases the version number.
	// Editing the issue will be blocked if the version reaches the maximum value: for robots 10100, for users 11100.
	Version int `json:"version"`
	// Date and time of the last comment added.
	LastCommentUpdatedAt string `json:"lastCommentUpdatedAt"`
	// Issue name.
	Summary string `json:"summary"`
	// No info: https://yandex.ru/support/tracker/ru/concepts/issues/search-issues
	Boards []Board `json:"boards"`
	// No info: https://yandex.ru/support/tracker/ru/concepts/issues/search-issues
	StatusStartTime string `json:"statusStartTime"`
	// An object with information about the parent issue.
	Parent IssueParent `json:"parent"`
	// An object containing information about the last employee who modified the issue.
	UpdatedBy IssueUpdatedBy `json:"updatedBy"`
	// Description of the issue.
	Description string `json:"description"`
	// An array of objects containing information about the sprint.
	Sprint IssueSprint `json:"sprint"`
	// An object with information about the issue type.
	Type IssueType `json:"type"`
	// Object with priority information.
	Priority IssuePriority `json:"priority"`
	// Date and time the issue was created.
	CreatedAt string `json:"createdAt"`
	// An array of objects containing information about the issue's observers.
	Followers IssueFollowers `json:"followers"`
	// An object containing information about the issue creator.
	CreatedBy IssueCreatedBy `json:"createdBy"`
	// An object with information about the issue performer.
	Assignee IssueAssignee `json:"assignee"`
	// Number of votes for the issue.
	Votes int `json:"votes"`
	// Number of comments with external message
	CommentWithExternalMessageCount int `json:"commentWithExternalMessageCount"`
	// Number of comments without external message
	CommentWithoutExternalMessageCount int `json:"commentWithoutExternalMessageCount"`
	// An object with information about a issue queue.
	Queue IssueQueue `json:"queue"`
	// Date and time the issue was updated.
	UpdatedAt string `json:"updatedAt"`
	// An object with information about the issue status.
	Status IssueStatus `json:"status"`
	// An object with information about the issue status type.
	StatusType IssueStatusType `json:"statusType"`
	// Flag of a favorite issue:
	// true — the user added the issue to favorites;
	// false — the issue was not added to favorites.
	Favorite bool `json:"favorite"`
	// An array of objects describing the attachments
	Attachments []Attachment `json:"attachments"`
	// An array of objects describing possible issue transitions
	Transitions []Transition `json:"transitions"`
}
