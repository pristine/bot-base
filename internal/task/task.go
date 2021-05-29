package task

// NOTE:
// there is a better way to handle tasks with interfaces

import (
	"errors"
	"github.com/EdwinJ0124/bot-base/internal/profile"
	"github.com/lithammer/shortuuid"
	"sync"
)

var (
	taskMutex = sync.RWMutex{}

	TaskDoesNotExistErr = errors.New("task does not exist")

	tasks = make(map[string]*Task)
)

// DoesTaskExist checks if a task exists
func DoesTaskExist(id string) bool {
	taskMutex.RLock()
	defer taskMutex.RUnlock()
	_, ok := tasks[id]
	return ok
}

// CreateTask creates a task
func CreateTask(taskType, site, product string) string {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	id := shortuuid.New()

	tasks[id] = &Task{
		Type:    taskType,
		Site:    site,
		Product: product,
	}

	return id
}

// RemoveTask removes a task
func RemoveTask(id string) error {
	if !DoesTaskExist(id) {
		return TaskDoesNotExistErr
	}

	taskMutex.Lock()
	defer taskMutex.Unlock()

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

	taskMutex.RLock()
	defer taskMutex.RUnlock()

	return tasks[id], nil
}

// SetProfileGroupToTask sets a profile group to a task
func SetProfileGroupToTask(taskId, profileGroupId string) error {
	if !DoesTaskExist(taskId) {
		return TaskDoesNotExistErr
	}

	if !profile.DoesProfileGroupExist(profileGroupId) {
		return profile.ProfileGroupDoesNotExistErr
	}

	taskMutex.Lock()
	defer taskMutex.Unlock()

	task := tasks[taskId]

	task.ProfileGroupID = profileGroupId

	return nil
}

// SetTaskToTaskGroup sets a task to a task group
func SetTaskToTaskGroup(taskId, taskGroupId string) error {
	if !DoesTaskExist(taskId) {
		return TaskDoesNotExistErr
	}

	if !DoesTaskGroupExist(taskGroupId) {
		return TaskGroupDoesNotExistErr
	}

	taskGroupMutex.Lock()
	defer taskGroupMutex.Unlock()

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

	taskMutex.Lock()
	defer taskMutex.Unlock()

	taskGroup := taskGroups[taskGroupId]

	delete(taskGroup.Tasks, taskId)

	return nil
}
