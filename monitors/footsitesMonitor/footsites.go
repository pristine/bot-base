package footsitesMonitor

import (
	"github.com/EdwinJ0124/bot-base/internal/monitor"
)

func Initialize() {
	monitorType := monitor.RegisterMonitorType("footsites")

	monitorType.SetFirstHandlerState(INITIALIZE)

	monitorType.AddHandlers(monitor.MonitorHandlerMap{
		INITIALIZE: initialize,
		GET_STOCK: getStock,
	})
}