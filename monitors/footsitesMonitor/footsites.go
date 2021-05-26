package footsitesMonitor

import (
	"github.com/EdwinJ0124/footsites/internal/monitor"
)

func Initialize() {
	monitorType := monitor.RegisterMonitorType("footsites")

	monitorType.SetFirstHandlerState(INITIALIZE)

	monitorType.AddHandlers(monitor.MonitorHandlerMap{
		INITIALIZE: initialize,
		GET_STOCK: getStock,
	})
}