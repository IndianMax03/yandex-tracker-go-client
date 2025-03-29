// The model package contains an entities for exchanging information with the Yandex Tracker API
package model

type IssueModifyStatusRequest struct {
	Resolution string `json:"resolution,omitempty"`
	Assignee   string `json:"assignee,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

type IssueModifyStatusResponse struct {
	Self         string         `json:"self"`
	TransitionID string         `json:"id"`
	To           NextTransition `json:"to"`
}

type NextTransition struct {
	ObjectBaseResponse
	Key string `json:"key"`
}

func newTransition() (req *IssueModifyStatusRequest) {
	req = &IssueModifyStatusRequest{}
	return
}

func NewReopenTransition() (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	transitionID = ReopenTransitionID
	return
}

func NewInProgressTransition(assignee string) (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	req.Assignee = assignee
	transitionID = InProgrssTransitionID
	return
}

func NewStopProgressTransition() (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	transitionID = StopProgressTransitionID
	return
}

func NewNeedInfoTransition(assignee string) (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	req.Assignee = assignee
	transitionID = NeedInfoTransitionID
	return
}

func NewProvideInfoTransition(assignee string) (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	req.Assignee = assignee
	transitionID = ProvideInfoTransitionID
	return
}

func NewCloseFixedTransition() (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	req.Resolution = Fixed
	transitionID = CloseTransitionID
	return
}

func NewCloseWontFixTransition() (req *IssueModifyStatusRequest, transitionID string) {
	req = newTransition()
	req.Resolution = WontFix
	transitionID = CloseTransitionID
	return
}
