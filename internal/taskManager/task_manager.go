package taskManager

import (
	"context"
	"github.com/EdwinJ0124/footsites/internal/task"
	"time"
)

func handleTaskState(taskState task.TaskState, taskType *task.TaskType, t *task.Task) task.TaskState {
	nextTaskHandler := taskType.GetHandler(taskState)

	nextNextTaskType := nextTaskHandler(t)

	return nextNextTaskType
}

// RunTask starts a task
func RunTask(t *task.Task) {
	t.Context, t.Cancel = context.WithCancel(context.Background())
	t.Active = true

	defer func() {
		if r := recover(); r != nil {
			// handle crash
		}
	}()

	if !task.DoesTaskTypeExist(t.Type) {
		return
	}

	taskType, err := task.GetTaskType(t.Type)

	if err != nil {
		t.Active = false
		return
	}

	hasHandlers := taskType.HasHandlers()

	if !hasHandlers {
		t.Active = false
		return
	}

	nextState := taskType.GetFirstHandlerState()

	// loop the task states
	for {
		nextState = handleTaskState(nextState, taskType, t)

		if nextState == task.DoneTaskState || t.Context.Err() != nil  {
			// you can report that the task stopped here
			t.Active = false
			break
		}

		time.Sleep(1 * time.Millisecond)
	}
}

// StopTask stops a task
func StopTask(t *task.Task) {
	t.Cancel()
}