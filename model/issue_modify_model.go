// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// IssueModifyRequest describes request to modify existing issue
type IssueModifyRequest struct {
	// Issue name.
	Summary string `json:"summary,omitempty"`
	// Parental issue.
	Parent ObjectBaseRequest `json:"parent,omitzero"`
	// Description of the issue.
	Description string `json:"description,omitempty"`
	// The type of markup displayed in the text.
	// If you use YFM markup in the issue description text, specify the md value.
	MarkupType string `json:"markupType,omitempty"`
	// An array of objects containing information about the sprint.
	Sprint SprintModifyRequest `json:"sprint,omitzero"`
	// An object with information about the issue type.
	Type ObjectBaseRequest `json:"type,omitzero"`
	// Object with priority information.
	Priority ObjectBaseRequest `json:"priority,omitzero"`
	// An array of objects containing information about the issue's observers.
	Followers ModifyFollowers `json:"followers,omitzero"`
	// IDs of temporary files to be added as attachments.
	AttachmentIds []string `json:"attachmentIds,omitempty"`
	// Identifiers of temporary files that will be added to the issue description.
	DescriptionAttachmentIds []string `json:"descriptionAttachmentIds,omitempty"`
	// An array of objects containing information about tags.
	Tags []string `json:"tags,omitempty"`
}

// ModifyFollowers describes request object to modify followers of existing issue
type ModifyFollowers struct {
	// List of people's IDs or logins
	Add []string `json:"add,omitempty"`
	// List of people's IDs or logins
	Remove []string `json:"remove,omitempty"`
}

// SprintModifyRequest describes sprint object to modify sprints of existing issue
type SprintModifyRequest struct {
	ID string `json:"id"`
}
