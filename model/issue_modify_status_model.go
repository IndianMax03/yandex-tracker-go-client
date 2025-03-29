// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

// IssueModifyStatusRequest describes request to modify issue status
type IssueModifyStatusRequest struct {
	// Issue  resolution
	Resolution string `json:"resolution,omitempty"`
	// Issue field available for modification during transition. List of keys: https://tracker.yandex.ru/admin/fields
	Assignee string `json:"assignee,omitempty"`
	// Commentary on the issue.
	Comment string `json:"comment,omitempty"`
}

// IssueModifyStatusResponse describes response of issue status modify contains next allowed issue transitions
type IssueModifyStatusResponse struct {
	// The address of the API resource that contains the transition information.
	Self string `json:"self"`
	// Transition identifier.
	TransitionID string `json:"id"`
	// A block with status information to which a issue can be transferred.
	To NextTransition `json:"to"`
}

// NextTransition describes next possible issue transitions.
type NextTransition struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

func newTransition() (req *IssueModifyStatusRequest) {
	req = &IssueModifyStatusRequest{}
	return
}

// NewReopenTransition instantiates transition to open
func NewReopenTransition() (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	transitionID = ReopenTransitionID
	return
}

// NewInProgressTransition instantiates transition to in progress
func NewInProgressTransition(assignee string) (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	req.Assignee = assignee
	transitionID = InProgrssTransitionID
	return
}

// NewStopProgressTransition instantiates transition to stop progress
func NewStopProgressTransition() (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	transitionID = StopProgressTransitionID
	return
}

// NewNeedInfoTransition instantiates transition to need info
func NewNeedInfoTransition(assignee string) (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	req.Assignee = assignee
	transitionID = NeedInfoTransitionID
	return
}

// NewProvideInfoTransition instantiates transition to provide info
func NewProvideInfoTransition(assignee string) (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	req.Assignee = assignee
	transitionID = ProvideInfoTransitionID
	return
}

// NewCloseFixedTransition instantiates transition to close (fixed)
func NewCloseFixedTransition() (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	req.Resolution = Fixed
	transitionID = CloseTransitionID
	return
}

// NewCloseWontFixTransition instantiates transition to close (Won't fix)
func NewCloseWontFixTransition() (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	req.Resolution = WontFix
	transitionID = CloseTransitionID
	return
}
