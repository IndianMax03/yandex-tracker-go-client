// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// ComponentRequest describes request to create a new component
type ComponentRequest struct {
	// Mandatory

	// Component name.
	Name string `json:"name"`
	// The key of the queue in which the component will be created.
	Queue string `json:"queue"`

	// Optional

	// Description of the component.
	Description string `json:"description,omitempty"`
	// Login of the component owner.
	Lead string `json:"lead,omitempty"`
	// Default executor flag:
	// true — assign the owner as the default executor;
	// false — do not assign the default executor.
	AssignAuto bool `json:"assignAuto,omitempty"`
}

// ComponentResponse describes component
type ComponentResponse struct {
	// The address of the API resource that contains the component's parameters.
	Self string `json:"self"`
	// Component identifier.
	ID int `json:"id"`
	// Component version. Each change to the component parameters increases the version number.
	Version int `json:"version"`
	// Component name.
	Name string `json:"name"`
	// An object with information about the queue to which the component has been added.
	Queue IssueQueue `json:"queue"`
	// The flag indicating whether the owner is assigned as the default executor.
	AssignAuto bool `json:"assignAuto"`
}
