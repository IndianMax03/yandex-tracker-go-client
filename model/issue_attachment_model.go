// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// AttachmentFileResponse describes response contains information about attachment.
type AttachmentFileResponse struct {
	// The address of the API resource that corresponds to the attached file.
	Self string `json:"self"`
	// Unique file identifier.
	ID string `json:"id"`
	// File name.
	Name string `json:"name"`
	// The address of the resource for downloading the file.
	Content string `json:"content"`
	// Resource address for downloading the preview thumbnail. Available only for graphic files.
	Thumbnail string `json:"thumbnail"`
	// An object containing information about the user who attached the file.
	CreatedBy CreatedBy `json:"createdBy"`
	// Date and time of file upload in the format: YYYY-MM-DDThh:mm:ss.sss±hhmm
	CreatedAt string `json:"createdAt"`
	// File type, for example:
	// text/plain — text file;
	// image/png — png image.
	Mimetype string `json:"mimetype"`
	// File size in bytes.
	Size int `json:"size"`
	// An object with file metadata.
	Metadata FileMetadata `json:"metadata"`
}

// FileMetadata describes object with file metadata.
type FileMetadata struct {
	// Image size in pixels
	Size string
}
