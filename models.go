// The client package provides methods, values and models for interacting with Yandex Tracker
package client

// Request to create a new issue
type IssueCreateRequest struct {
	// Mandatory

	// Issue name.
	Summary string `json:"summary"`
	// Queue in which the issue should be created.
	Queue Queue `json:"queue"`

	// Optional

	// Parent issue relative to current.
	Parent any `json:"parent,omitempty"`
	// Description of the issue.
	Description string `json:"description,omitempty"`
	// The type of markup displayed in the text.
	// If you use YFM markup in the issue description text, specify the md value.
	MarkupType string `json:"markupType,omitempty"`
	// An array of objects containing information about the sprint.
	Sprint []any `json:"sprint,omitempty"`
	// An object with information about the issue type.
	Type any `json:"type,omitempty"`
	// Object with priority information.
	Priority any `json:"priority,omitempty"`
	// An array of objects containing information about the issue's observers.
	Followers []any `json:"followers,omitempty"`
	// An object with information about the issue performer.
	Assignee any `json:"assignee,omitempty"`
	// The ID or login of the issue author
	Author any `json:"author,omitempty"`
	// A field with a unique value that prevents the creation of duplicate issues.
	// If you try to create a issue again with the same value for this parameter, a duplicate will not be created,
	// and the response will contain an error with the code 409.
	Unique string `json:"unique,omitempty"`
	// IDs of temporary files to be added as attachments.
	AttachmentIds []string `json:"attachmentIds,omitempty"`
	// IDs of temporary files that will be added to the issue description.
	DescriptionAttachmentIds []string `json:"descriptionAttachmentIds,omitempty"`
	// An array of objects containing information about tags.
	Tags []string `json:"tags,omitempty"`
}

// Newly created issue
type IssueCreateResponse struct {
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
	// Number of comments without external message
	CommentWithExternalMessageCount int `json:"commentWithExternalMessageCount"`
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
}

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

type Queue struct {
	ID  string `json:"id"`
	Key string `json:"key"`
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
