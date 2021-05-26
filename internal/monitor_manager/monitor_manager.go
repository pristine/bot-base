package monitormngr

import (
	"context"
	"github.com/EdwinJ0124/bot-base/internal/monitor"
	"time"
)

func handleMonitorState(monitorState monitor.MonitorState, monitorType *monitor.MonitorType, t *monitor.Monitor) monitor.MonitorState {
	nextHandler, err := monitorType.GetHandler(monitorState)

	if err != nil {
		return monitor.ErrorMonitorState
	}

	nextNextType := nextHandler(t)

	return nextNextType
}

// RunMonitor starts a monitor task
func RunMonitor(m *monitor.Monitor) {
	m.Context, m.Cancel = context.WithCancel(context.Background())
	m.Active = true

	defer func() {
		if r := recover(); r != nil {
			// handle crash
		}
	}()

	if !monitor.DoesMonitorTypeExist(m.Type) {
		return
	}

	monitorType, err := monitor.GetMonitorType(m.Type)

	if err != nil {
		m.Active = false
		return
	}

	hasHandlers := monitorType.HasHandlers()

	if !hasHandlers {
		m.Active = false
		return
	}

	nextState := monitorType.GetFirstHandlerState()

	if len(nextState) == 0 {
		m.Active = false
		return
	}

	// loop the moniitor states
	for {
		nextState = handleMonitorState(nextState, monitorType, m)

		if nextState == monitor.DoneMonitorState || m.Context.Err() != nil {
			// you can report that the monitor stopped here
			m.Active = false
			break
		} else if nextState == monitor.ErrorMonitorState {
			// report errors
			m.Active = false
			break
		}

		time.Sleep(1 * time.Millisecond)
	}
}

// StopMonitor stops a monitor task
func StopMonitor(m *monitor.Monitor) {
	m.Cancel()
}
