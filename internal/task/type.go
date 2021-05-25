package task

import (
	"errors"
	"github.com/iancoleman/orderedmap"
)

type TaskState string

type TaskType struct {
	handlers *orderedmap.OrderedMap
}

var (
	DoneTaskState TaskState = "done"

	TaskTypeDoesNotExistErr = errors.New("task type does not exist")

	taskTypes = make(map[string]*TaskType)
)

// RegisterTaskType register task type
func RegisterTaskType(taskType string) {
	taskTypes[taskType] = &TaskType{
		handlers: orderedmap.New(),
	}
}

// DoesTaskTypeExist check if task type exists
func DoesTaskTypeExist(taskType string) bool {
	_, ok := taskTypes[taskType]
	return ok
}

// GetTaskType gets a task type
func GetTaskType(taskType string) (*TaskType, error) {
	if !DoesTaskTypeExist(taskType) {
		return &TaskType{}, TaskTypeDoesNotExistErr
	}

	return taskTypes[taskType], nil
}

// HasHandlers check if there are handlers
func (t *TaskType) HasHandlers() bool {
	handlerIds := t.handlers.Keys()
	return len(handlerIds) > 0
}

// AddHandler adds a handler to the task type
func (t *TaskType) AddHandler(handlerName TaskState, handler func(*Task)TaskState) {
	t.handlers.Set(string(handlerName), handler)
}

// AddHandlers adds multiple handles to a task type
func (t *TaskType) AddHandlers(handlers map[TaskState]func(*Task)TaskState) {
	for handlerName, handler := range handlers {
		t.AddHandler(handlerName, handler)
	}
}

// GetHandler gets a handler by handler name
func (t *TaskType) GetHandler(handlerName TaskState) func(*Task)TaskState {
	handler, _ := t.handlers.Get(string(handlerName))
	return handler.(func(*Task)TaskState)
}

// GetFirstHandlerState gets the first handler state
func (t *TaskType) GetFirstHandlerState() TaskState {
	handlerIds := t.handlers.Keys()
	return TaskState(handlerIds[0])
}