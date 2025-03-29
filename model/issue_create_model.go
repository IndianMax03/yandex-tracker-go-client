// The model package contains an entities for exchanging information with the Yandex Tracker API
package model

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
