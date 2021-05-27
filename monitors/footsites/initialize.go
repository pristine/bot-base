package footsitesmntr

import "github.com/EdwinJ0124/bot-base/internal/monitor"

func initialize(m *monitor.Monitor, internal *FootsitesMonitorInternal) monitor.MonitorState {
	switch m.Params["site"] {
	case "footlocker":
		internal.Host = "www.footlocker.com"
	case "footaction":
		internal.Host = "www.footaction.com"
	case "eastbay":
		internal.Host = "www.eastbay.com"
	case "champssports":
		internal.Host = "www.champssports.com"
	case "footlockerca":
		internal.Host = "www.footlocker.ca"
	default:
		internal.Host = "www.footlocker.com"
	}

	return monitor.DoneMonitorState
}
