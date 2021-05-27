package footsitesmntr

import "github.com/EdwinJ0124/bot-base/internal/monitor"

func getStock(m *monitor.Monitor, internal *FootsitesMonitorInternal) monitor.MonitorState {

	// send a request to

	return HandleStockResponse(m, internal)
}

func HandleStockResponse(m *monitor.Monitor, internal *FootsitesMonitorInternal) monitor.MonitorState {
	// once in stock, you can simple notify all tasks associated with this monitor by doing

	//m.NotifyTasks(&FootsitesMonitorData{
	//	// necessary data needed to be sent
	//})

	return GET_STOCK
}
