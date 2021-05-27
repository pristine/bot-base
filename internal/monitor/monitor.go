package monitor

import (
	"errors"
	"github.com/EdwinJ0124/bot-base/internal/task"
	"github.com/lithammer/shortuuid"
	"sync"
)

var (
	monitorMutex = sync.RWMutex{}

	MonitorNotInTaskGroupErr = errors.New("monitor not in any task group")
	MonitorDoesNotExistErr   = errors.New("monitor does not exist")

	monitors = make(map[string]*Monitor)
)

// DoesMonitorExist checks if a monitor exists
func DoesMonitorExist(id string) bool {
	monitorMutex.RLock()
	defer monitorMutex.RUnlock()
	_, ok := monitors[id]
	return ok
}

// CreateMonitor creates a monitor
func CreateMonitor(input string) string {
	monitorMutex.Lock()
	defer monitorMutex.Unlock()
	id := shortuuid.New()

	monitors[id] = &Monitor{
		ID:    id,
		Input: input,
	}

	return id
}

// RemoveMonitor removes a monitor
func RemoveMonitor(id string) error {
	if !DoesMonitorExist(id) {
		return MonitorDoesNotExistErr
	}

	monitorMutex.Lock()
	defer monitorMutex.Unlock()

	monitor := monitors[id]
	monitor.Cancel()

	delete(monitors, id)

	return nil
}

// GetMonitor gets a monitor
func GetMonitor(id string) (*Monitor, error) {
	if !DoesMonitorExist(id) {
		return &Monitor{}, MonitorDoesNotExistErr
	}

	monitorMutex.RLock()
	defer monitorMutex.RUnlock()

	return monitors[id], nil
}

// SetMonitorToTaskGroup sets a task group to a monitor
func SetMonitorToTaskGroup(monitorId, taskGroupId string) error {
	if !DoesMonitorExist(monitorId) {
		return MonitorDoesNotExistErr
	}

	if !task.DoesTaskGroupExist(taskGroupId) {
		return task.TaskGroupDoesNotExistErr
	}

	taskGroup, _ := task.GetTaskGroup(monitorId)
	taskGroup.Monitors[monitorId] = true

	return nil
}

// NotifyTasks notifies tasks
func (m *Monitor) NotifyTasks(monitorData interface{}) error {
	associatedTaskGroup, err := m.getAssociatedTaskGroup()

	if err != nil {
		return err
	}

	taskIds := associatedTaskGroup.GetAllTaskIDs()

	for _, id := range taskIds {
		task, _ := task.GetTask(id)

		task.MonitorData <- monitorData
	}

	return nil
}

func (m *Monitor) getAssociatedTaskGroup() (*task.TaskGroup, error) {
	taskGroupIds := task.GetAllTaskGroupIDs()

	for _, taskGroupId := range taskGroupIds {
		taskGroup, _ := task.GetTaskGroup(taskGroupId)

		if taskGroup.Monitors[m.ID] {
			return taskGroup, nil
		}
	}

	return &task.TaskGroup{}, MonitorNotInTaskGroupErr
}
