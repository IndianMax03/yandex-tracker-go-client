// Model package contains an entities for exchanging information with the Yandex Tracker API
package model

// IssueModifyRequest describes request to modify existing issue
type IssueModifyRequest struct {
	Summary                  string              `json:"summary,omitempty"`
	Parent                   ObjectBaseRequest   `json:"parent,omitzero"`
	Description              string              `json:"description,omitempty"`
	MarkupType               string              `json:"markupType,omitempty"`
	Sprint                   SprintModifyRequest `json:"sprint,omitzero"`
	Type                     ObjectBaseRequest   `json:"type,omitzero"`
	Priority                 ObjectBaseRequest   `json:"priority,omitzero"`
	Followers                ModifyFollowers     `json:"followers,omitzero"`
	AttachmentIds            []string            `json:"attachmentIds,omitempty"`
	DescriptionAttachmentIds []string            `json:"descriptionAttachmentIds,omitempty"`
	Tags                     []string            `json:"tags,omitempty"`
}

// ModifyFollowers describes request object to modify followers of existing issue
type ModifyFollowers struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

// SprintModifyRequest describes sprint object to modify sprints of existing issue
type SprintModifyRequest struct {
	ID string `json:"id"`
}
