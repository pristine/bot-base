package footsites

import "github.com/EdwinJ0124/footsites/internal/task"

func initialize(t *task.Task) task.TaskState {
	t.Internal = &footsites{}



	return task.DoneTaskState
}
