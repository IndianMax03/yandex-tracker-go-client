// The client package provides methods, values and urls for interacting with Yandex Tracker
package client

import (
	"fmt"
	"strconv"

	model "github.com/IndianMax03/yandex-tracker-go-client/model"
)

const (
	TrivialPriority = iota + 1
	MinorPriority
	NormalPriority
	CriticalPriority
	BlockerPriority
)

func GetPriority(priority int) (*model.ObjectBaseRequest, error) {
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
	return &model.ObjectBaseRequest{
		ID:  strconv.Itoa(priority),
		Key: priorityName,
	}, nil
}
