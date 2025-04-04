// Package model contains an entities for exchanging information with the Yandex Tracker API
package model

import (
	"fmt"
	"strconv"
)

// Issue priority
const (
	TrivialPriority = iota + 1
	MinorPriority
	NormalPriority
	CriticalPriority
	BlockerPriority
)

// Issue resolution
const (
	Fixed         = "fixed"
	WontFix       = "wontFix"
	CantReproduce = "cantReproduce"
	Duplicate     = "duplicate"
	Later         = "later"
	Overfulfilled = "overfulfilled"
	Successful    = "successful"
	DontDo        = "dontDo"
)

// Comment expand
const (
	ExpandAttachments = "attachments"
	ExpandHTML        = "html"
	ExpandAll         = "all"
	ExpandNone        = ""
)

// Issue Transition IDs
const (
	InProgrssTransitionID    = "start_progress"
	StopProgressTransitionID = "stop_progress"
	NeedInfoTransitionID     = "need_info"
	ProvideInfoTransitionID  = "provide_info"
	CloseTransitionID        = "close"
	ReopenTransitionID       = "reopen"
)

// GetPriority maps priority value to priority id & key
func GetPriority(priority int) (*ObjectBaseRequest, error) {
	priorityName := ""
	if priority == TrivialPriority {
		priorityName = "trivial"
	} else if priority == MinorPriority {
		priorityName = "minor"
	} else if priority == NormalPriority {
		priorityName = "normal"
	} else if priority == CriticalPriority {
		priorityName = "critical"
	} else if priority == BlockerPriority {
		priorityName = "blocker"
	} else {
		return nil, fmt.Errorf("unknown priority: %v", priority)
	}
	return &ObjectBaseRequest{
		ID:  strconv.Itoa(priority),
		Key: priorityName,
	}, nil
}
