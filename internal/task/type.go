package task

import (
	"errors"
	"github.com/iancoleman/orderedmap"
)

type TaskState string
type TaskHandlerMap map[TaskState]func(*Task)TaskState

type TaskType struct {
	firstHandlerState TaskState
	handlers *orderedmap.OrderedMap
}

var (
	DoneTaskState TaskState = "done"
	ErrorTaskState TaskState = "error"

	TaskTypeDoesNotExistErr = errors.New("task type does not exist")
	TaskHandlerDoesNotExistErr = errors.New("task handler does not exist")

	taskTypes = make(map[string]*TaskType)
)

// RegisterTaskType register task type
func RegisterTaskType(taskType string) *TaskType {
	taskTypes[taskType] = &TaskType{
		handlers: orderedmap.New(),
	}
	return taskTypes[taskType]
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
func (t *TaskType) AddHandlers(handlers TaskHandlerMap) {
	for handlerName, handler := range handlers {
		t.AddHandler(handlerName, handler)
	}
}

// GetHandler gets a handler by handler name
func (t *TaskType) GetHandler(handlerName TaskState) (func(*Task)TaskState, error) {
	handler, ok := t.handlers.Get(string(handlerName))

	if !ok {
		return nil, TaskHandlerDoesNotExistErr
	}

	return handler.(func(*Task)TaskState), nil
}

// GetFirstHandlerState gets the first handler state
func (t *TaskType) GetFirstHandlerState() TaskState {
	return t.firstHandlerState
}

func (t *TaskType) SetFirstHandlerState(firstHandlerState TaskState) {
	t.firstHandlerState = firstHandlerState
}