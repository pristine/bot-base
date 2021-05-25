package footsites

import "github.com/EdwinJ0124/footsites/internal/task"

func Initialize() {
	taskType := task.RegisterTaskType("footsites")

	taskType.SetFirstHandlerState(INITIALIZE)

	taskType.AddHandlers(task.TaskHandlerMap{
		INITIALIZE: initialize,
	})
}