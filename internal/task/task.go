package task

import (
	"context"
	"errors"
	"github.com/EdwinJ0124/footsites/pkg/hclient"
	"github.com/lithammer/shortuuid"
)

type Task struct {
	ID             string             `json:"id"`
	Site           string             `json:"site"`
	ProfileGroupID string             `json:"profileGroupID"`
	ProxyListID    string             `json:"proxyListID"`
	Context        context.Context    `json:"-"`
	Cancel         context.CancelFunc `json:"-"`
	Internal       interface{}        `json:"-"`
	Client		   *hclient.Client     `json:"-"`
}

var (
	TaskDoesNotExistErr = errors.New("task does not exist")

	tasks = make(map[string]*Task)
)

// DoesTaskExist checks if a task exists
func DoesTaskExist(id string) bool {
	_, ok := tasks[id]
	return ok
}

// CreateTask creates a task
func CreateTask(site string) string {
	id := shortuuid.New()

	tasks[id] = &Task{
		Site: site,
	}

	return id
}

func AssignProfileGroupToTask(taskId, profileGroupId string) error {
	if !DoesTaskExist(taskId) {
		return TaskDoesNotExistErr
	}


}