package footsitesMonitor

import "github.com/EdwinJ0124/footsites/internal/monitor"

func getStock(m *monitor.Monitor) monitor.MonitorState {

	// send a request to


	return HandleStockResponse(m)
}

func HandleStockResponse(m *monitor.Monitor) monitor.MonitorState {
	// once in stock, you can simple notify all tasks associated with this monitor by doing

	//m.NotifyTasks(&FootsitesMonitorData{
	//	// necessary data needed to be sent
	//})

	return GET_STOCK
}