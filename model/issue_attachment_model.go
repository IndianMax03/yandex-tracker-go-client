// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// AttachmentFileResponse describes response contains information about attachment.
type AttachmentFileResponse struct {
	// The address of the API resource that corresponds to the attached file.
	Self string
	// Unique file identifier.
	ID string
	// File name.
	Name string
	// The address of the resource for downloading the file.
	Content string
	// Resource address for downloading the preview thumbnail. Available only for graphic files.
	Thumbnail string
	// An object containing information about the user who attached the file.
	CreatedBy CreatedBy
	// Date and time of file upload in the format: YYYY-MM-DDThh:mm:ss.sss±hhmm
	CreatedAt string
	// File type, for example:
	// text/plain — text file;
	// image/png — png image.
	Mimetype string
	// File size in bytes.
	Size int
	// An object with file metadata.
	Metadata FileMetadata
}

// FileMetadata describes object with file metadata.
type FileMetadata struct {
	// Image size in pixels
	Size string
}
