// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// CommentRequest describes request to add comment on the issue.
type CommentRequest struct {
	// Mandatory

	// Commentary on the issue.
	Text string `json:"text"`

	// Optional

	// IDs of temporary files to be added as attachments.
	AttachmentIds []string `json:"attachmentIds,omitempty"`
	// IDs or logins of the invited users.
	Summonees []string `json:"summonees,omitempty"`
	// List of mailings called in comments.
	MaillistSummonees []string `json:"maillistSummonees,omitempty"`
	// The type of markup displayed in the text.
	// If you use YFM markup in the issue description text, specify the md value.
	MarkupType string `json:"markupType,omitempty"`
}

// CommentResponse describes response after adding comment on the issue.
type CommentResponse struct {
	// Link to comment.
	Self string `json:"self"`
	// Comment ID.
	ID int `json:"id"`
	// Comment identifier in string format.
	LongID string `json:"longId"`
	// Commentary on the issue.
	Text string `json:"text"`
	// HTML markup of the comment.
	TextHTML string `json:"textHtml"`
	// Attachments to the comment.
	Attachments []Attachment `json:"attachments"`
	// Block with information about the user who added the comment.
	CreatedBy CreatedBy `json:"createdBy"`
	// Block with information about the user who last changed the comment.
	UpdatedBy UpdatedBy `json:"updatedBy"`
	// Date and time the comment was created.
	CreatedAt string `json:"createdAt"`
	// Date and time the comment was updated.
	UpdatedAt string `json:"updatedAt"`
	// Block with information about users who are invited to comment.
	Summonees []Summonees `json:"summonees"`
	// Block with information about mailings that are called in comments.
	Maillistsummonees []Maillistsummonees `json:"maillistsummonees"`
	// Comment version. Each change to the comment increments the version number.
	Version int `json:"version"`
	// Comment type:
	// standard — sent via the Tracker interface;
	// incoming — created from an incoming email;
	// outcoming — created from an outgoing email.
	Type string `json:"type"`
	// Method of adding a comment:
	// internal — via the Tracker interface;
	// email — via letter.
	Transport string `json:"transport"`
}

// Summonees is a users who are invited.
type Summonees ObjectBaseResponse

// Maillistsummonees is a mailings that are called.
type Maillistsummonees ObjectBaseResponse
