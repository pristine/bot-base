package footsites

import "github.com/EdwinJ0124/bot-base/internal/task"

func submitOrder(t *task.Task) task.TaskState {



	return handleSubmitOrderResponse(t)
}

func handleSubmitOrderResponse(t *task.Task) task.TaskState {

	return task.DoneTaskState
}