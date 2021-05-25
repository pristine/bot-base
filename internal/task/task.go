package task

import (
	"context"
	"errors"
	"github.com/EdwinJ0124/footsites/internal/profile"
	"github.com/EdwinJ0124/footsites/third_party/hclient"
	"github.com/lithammer/shortuuid"
)

type Task struct {
	ID             string             `json:"id"`
	Params		   map[string]string   `json:"params"`
	Type           string             `json:"type"`
	ProfileGroupID string             `json:"profileGroupID"`
	ProxyListID    string             `json:"proxyListID"`
	Context        context.Context    `json:"-"`
	Cancel         context.CancelFunc `json:"-"`
	Internal       interface{}        `json:"-"`
	Client		   *hclient.Client    `json:"-"`
	Active         bool               `json:"-"`
	MonitorData	   interface{}		  `json:"-"`
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
func CreateTask(taskType string, params map[string]string) string {
	id := shortuuid.New()

	tasks[id] = &Task{
		Type: taskType,
		Params: params,
	}

	return id
}

// RemoveTask removes a task
func RemoveTask(id string) error {
	if !DoesTaskExist(id) {
		return TaskDoesNotExistErr
	}

	// stop the task if active
	task := tasks[id]
	task.Cancel()

	delete(tasks, id)

	return nil
}

// GetTask gets a task
func GetTask(id string) (*Task, error) {
	if !DoesTaskExist(id) {
		return &Task{}, TaskDoesNotExistErr
	}

	return tasks[id], nil
}

// AssignProfileGroupToTask assigns a profile group to a task
func AssignProfileGroupToTask(taskId, profileGroupId string) error {
	if !DoesTaskExist(taskId) {
		return TaskDoesNotExistErr
	}

	if !profile.DoesProfileGroupExist(profileGroupId) {
		return profile.ProfileGroupDoesNotExistErr
	}

	task := tasks[taskId]

	task.ProfileGroupID = profileGroupId

	return nil
}

// AssignTaskToTaskGroup assigns a task to a task group
func AssignTaskToTaskGroup(taskId, taskGroupId string) error {
	if !DoesTaskExist(taskId) {
		return TaskDoesNotExistErr
	}

	if !DoesTaskGroupExist(taskGroupId) {
		return TaskGroupDoesNotExistErr
	}

	taskGroup := taskGroups[taskGroupId]

	taskGroup.Tasks[taskId] = true

	return nil
}

// RemoveTaskFromTaskGroup removes a task from a task group
func RemoveTaskFromTaskGroup(taskId, taskGroupId string) error {
	if !DoesTaskExist(taskId) {
		return TaskDoesNotExistErr
	}

	if !DoesTaskGroupExist(taskGroupId) {
		return TaskGroupDoesNotExistErr
	}

	taskGroup := taskGroups[taskGroupId]

	delete(taskGroup.Tasks, taskId)

	return nil
}