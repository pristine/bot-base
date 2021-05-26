package monitor

import (
	"errors"
	"github.com/iancoleman/orderedmap"
)

type MonitorState string
type MonitorHandlerMap map[MonitorState]func(*Monitor) MonitorState

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
		handlers: orderedmap.New(),
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
	handlerIds := t.handlers.Keys()
	return len(handlerIds) > 0
}

// AddHandler adds a handler to the monitor type
func (t *MonitorType) AddHandler(handlerName MonitorState, handler func(*Monitor) MonitorState) {
	t.handlers.Set(string(handlerName), handler)
}

// AddHandlers adds multiple handles to a monitor type
func (t *MonitorType) AddHandlers(handlers MonitorHandlerMap) {
	for handlerName, handler := range handlers {
		t.AddHandler(handlerName, handler)
	}
}

// GetHandler gets a handler by handler name
func (t *MonitorType) GetHandler(handlerName MonitorState) (func(*Monitor) MonitorState, error) {
	handler, ok := t.handlers.Get(string(handlerName))

	if !ok {
		return nil, MonitorHandlerDoesNotExistErr
	}

	return handler.(func(*Monitor) MonitorState), nil
}

// GetFirstHandlerState gets the first handler state
func (t *MonitorType) GetFirstHandlerState() MonitorState {
	return t.firstHandlerState
}

func (t *MonitorType) SetFirstHandlerState(firstHandlerState MonitorState) {
	t.firstHandlerState = firstHandlerState
}
