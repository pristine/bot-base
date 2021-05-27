package monitor

import (
	"errors"
	"reflect"
)

var (
	DoneMonitorState  MonitorState = "done"
	ErrorMonitorState MonitorState = "error"

	MonitorTypeDoesNotExistErr    = errors.New("monitor type does not exist")
	MonitorHandlerDoesNotExistErr = errors.New("monitor handler does not exist")

	monitorTypes = make(map[string]*MonitorType)
)

// RegisterMonitorType register monitor type
func RegisterMonitorType(monitorType string) *MonitorType {
	monitorTypes[monitorType] = &MonitorType{
		handlers: make(MonitorReflectMap),
	}

	return monitorTypes[monitorType]
}

// DoesMonitorTypeExist check if monitor type exists
func DoesMonitorTypeExist(monitorType string) bool {
	_, ok := monitorTypes[monitorType]
	return ok
}

// GetMonitorType gets a task type
func GetMonitorType(monitorType string) (*MonitorType, error) {
	if !DoesMonitorTypeExist(monitorType) {
		return &MonitorType{}, MonitorTypeDoesNotExistErr
	}
	return monitorTypes[monitorType], nil
}

// HasHandlers check if there are handlers
func (t *MonitorType) HasHandlers() bool {
	return len(t.handlers) > 0
}

func (t *MonitorType) addHandler(handlerName MonitorState, handler interface{}) {
	t.handlers[string(handlerName)] = reflect.ValueOf(handler)
}

// AddHandlers adds multiple handles to a monitor type
func (t *MonitorType) AddHandlers(handlers MonitorHandlerMap) {
	for handlerName, handler := range handlers {
		if t.internalType == nil {
			handleTypes := reflect.TypeOf(handler)
			// func (t *task.Task, internal *SiteInternal) task.TaskState

			// we want the second one because the first one (0 index) will be task.Task type
			handleType := handleTypes.In(1)

			t.internalType = handleType
		}

		t.addHandler(handlerName, handler)
	}
}

// GetHandler gets a handler by handler name
func (t *MonitorType) GetHandler(handlerName MonitorState) (reflect.Value, error) {
	handler, ok := t.handlers[string(handlerName)]

	if !ok {
		return reflect.Value{}, MonitorHandlerDoesNotExistErr
	}

	return handler, nil
}

// GetFirstHandlerState gets the first handler state
func (t *MonitorType) GetFirstHandlerState() MonitorState {
	return t.firstHandlerState
}

// SetFirstHandlerState sets the first handler state
func (t *MonitorType) SetFirstHandlerState(firstHandlerState MonitorState) {
	t.firstHandlerState = firstHandlerState
}

// GetInternalType gets the internal type
func (t *MonitorType) GetInternalType() reflect.Type {
	return t.internalType
}
