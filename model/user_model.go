// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// MyselfResponse describes an object contains information about current user.
type MyselfResponse struct {
	// The address of the API resource that contains information about the user account.
	Self string `json:"self"`
	// Unique identifier of the user account in Tracker.
	UID int `json:"uid"`
	// User login.
	Login string `json:"login"`
	// Unique identifier of the user account in Tracker.
	TrackerUID int `json:"trackerUid"`
	// Unique identifier of the user account in the Yandex 360 for Business organization and Yandex ID.
	PassportUID int `json:"passportUid"`
	// Unique user identifier in Yandex Cloud Organization.
	CloudUID string `json:"cloudUid"`
	// Username.
	FirstName string `json:"firstName"`
	// User's last name.
	LastName string `json:"lastName"`
	// Display name of the user.
	Display string `json:"display"`
	// User's email.
	Email string `json:"email"`
	// Indicates whether the user has full access to Tracker:
	// true — full access;
	// false — read only.
	HasLicense bool `json:"hasLicense"`
	// User status in the organization:
	// true — the user has been removed from the organization;
	// false — a current employee of the organization.
	Dismissed bool `json:"dismissed"`
	// Flag for forcibly disabling notifications for the user:
	// true — notifications are disabled;
	// false — notifications are enabled.
	DisableNotifications bool `json:"disableNotifications"`
	// Date and time of the user's first authorization in the format YYYY-MM-DDThh:mm:ss.sss±hhmm.
	FirstLoginDate string `json:"firstLoginDate"`
	// Date and time of the user's last authorization in the format YYYY-MM-DDThh:mm:ss.sss±hhmm.
	LastLoginDate string `json:"lastLoginDate"`
	// Method for adding a user:
	// true — via email invitation;
	// false — by other means.
	WelcomeMailSent bool `json:"welcomeMailSent"`
}
