package task

import (
	"errors"
	"github.com/lithammer/shortuuid"
	"sync"
)

var (
	taskGroupMutex = sync.RWMutex{}

	TaskGroupDoesNotExistErr = errors.New("task group does not exist")

	taskGroups = make(map[string]*TaskGroup)
)

// DoesTaskGroupExist determines if a task group is present
func DoesTaskGroupExist(id string) bool {
	taskGroupMutex.RLock()
	defer taskGroupMutex.RUnlock()
	_, ok := taskGroups[id]
	return ok
}

// CreateTaskGroup creates a new task group
func CreateTaskGroup(name string) string {
	taskGroupMutex.Lock()
	defer taskGroupMutex.Unlock()
	id := shortuuid.New()

	taskGroups[id] = &TaskGroup{
		Name:  name,
		ID:    id,
		Tasks: make(map[string]bool),
	}

	return id
}

// RemoveTaskGroup removes a specified task group
func RemoveTaskGroup(id string) error {
	if !DoesTaskGroupExist(id) {
		return TaskGroupDoesNotExistErr
	}

	taskGroupMutex.Lock()
	defer taskGroupMutex.Unlock()

	delete(taskGroups, id)

	return nil
}

// GetTaskGroup gets a task group from a specified id
func GetTaskGroup(id string) (*TaskGroup, error) {
	if !DoesTaskGroupExist(id) {
		return &TaskGroup{}, TaskGroupDoesNotExistErr
	}

	taskGroupMutex.RLock()
	defer taskGroupMutex.RUnlock()

	return taskGroups[id], nil
}

// GetTaskIDs gets all task ids of a specified group
func GetTaskIDs(id string) ([]string, error) {
	if !DoesTaskGroupExist(id) {
		return []string{}, TaskGroupDoesNotExistErr
	}

	taskGroupMutex.RLock()
	defer taskGroupMutex.RUnlock()

	ids := make([]string, 0)

	taskGroup := taskGroups[id]

	for id := range taskGroup.Tasks {
		ids = append(ids, id)
	}

	return ids, nil
}

// GetAllTaskGroupIDs gets all task group ids
func GetAllTaskGroupIDs() []string {
	taskGroupMutex.RLock()
	defer taskGroupMutex.RUnlock()

	ids := make([]string, 0)

	for id := range taskGroups {
		ids = append(ids, id)
	}

	return ids
}

// GetAllTaskIDs gets all task ids
func (t *TaskGroup) GetAllTaskIDs() []string {
	ids := make([]string, 0)

	for id := range t.Tasks {
		ids = append(ids, id)
	}

	return ids
}

// HasMonitors checks if task groups has monitors
func (t *TaskGroup) HasMonitors() bool {
	return len(t.Monitors) > 0
}

// GetAllMonitorIDs gets all monitor ids
func (t *TaskGroup) GetAllMonitorIDs() []string {
	ids := make([]string, 0)

	for id := range t.Monitors {
		ids = append(ids, id)
	}

	return ids
}
