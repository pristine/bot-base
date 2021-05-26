package footsites

import "github.com/EdwinJ0124/bot-base/internal/task"

var (
	INITIALIZE task.TaskState = "initialize"
	GET_SESSION task.TaskState = "get_session"
	WAIT_FOR_MONITOR task.TaskState = "wait_for_monitor"
)