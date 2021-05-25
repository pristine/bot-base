package monitor

import (
	"context"
	"errors"
	"github.com/EdwinJ0124/footsites/internal/task"
	"github.com/EdwinJ0124/footsites/third_party/hclient"
	"github.com/lithammer/shortuuid"
)

type Monitor struct {
	ID string 						  `json:"id"`
	Type string						  `json:"type"`
	Input string 					  `json:"input"`
	ProxyListID    string             `json:"proxyListID"`
	Internal       interface{}        `json:"-"`
	Context        context.Context    `json:"-"`
	Cancel         context.CancelFunc `json:"-"`
	Client *hclient.Client 			  `json:"-"`
	Active	bool					  `json:"-"`
}

var(
	MonitorNotInTaskGroupErr = errors.New("monitor not in any task group")
	MonitorDoesNotExistErr = errors.New("monitor does not exist")

	monitors = make(map[string]*Monitor)
)

// DoesMonitorExist checks if a monitor exists
func DoesMonitorExist(id string) bool {
	_, ok := monitors[id]
	return ok
}

// CreateMonitor creates a monitor
func CreateMonitor(input string) string {
	id := shortuuid.New()

	monitors[id] = &Monitor{
		ID: id,
		Input: input,
	}

	return id
}

// RemoveMonitor removes a monitor
func RemoveMonitor(id string) error {
	if !DoesMonitorExist(id) {
		return MonitorDoesNotExistErr
	}

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
	return monitors[id], nil
}

// AssignMonitorToTaskGroup assigns a task group to a monitor
func AssignMonitorToTaskGroup(monitorId, taskGroupId string) error {
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
func (m *Monitor) NotifyTasks(id string, monitorData interface{}) error {
	if !DoesMonitorExist(id) {
		return MonitorDoesNotExistErr
	}

	associatedTaskGroup, err := m.GetAssociatedTaskGroup()

	if err != nil {
		return err
	}

	taskIds := associatedTaskGroup.GetAllTaskIDs()

	for _, id := range taskIds {
		task, _ := task.GetTask(id)

		task.MonitorData = monitorData
	}

	return nil
}

func (m *Monitor) GetAssociatedTaskGroup() (*task.TaskGroup, error) {
	taskGroupIds := task.GetAllTaskGroupIDs()

	for _, taskGroupId := range taskGroupIds {
		taskGroup, _ := task.GetTaskGroup(taskGroupId)

		if taskGroup.Monitors[m.ID] {
			return taskGroup, nil
		}
	}

	return &task.TaskGroup{}, MonitorNotInTaskGroupErr
}