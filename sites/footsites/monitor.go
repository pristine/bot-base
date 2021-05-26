package footsites

import (
	"github.com/EdwinJ0124/footsites/internal/task"
)

func waitForMonitor(t *task.Task) task.TaskState {
	// wait for data to come from montitor

	// data := <- t.MonitorData

	// monitorData := data.(*footsitesMonitor.FootsitesMonitorData)

	// handle data

	return task.DoneTaskState
}