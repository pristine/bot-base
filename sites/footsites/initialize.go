package footsites

import (
	"github.com/EdwinJ0124/bot-base/internal/proxy"
	"github.com/EdwinJ0124/bot-base/internal/task"
	"github.com/EdwinJ0124/bot-base/third_party/hclient"
)

// NOTE:
// there are better ways to do it via a new task system, but this works too

func initialize(t *task.Task) task.TaskState {
	t.Internal = &footsites{}

	internal := t.Internal.(*footsites)

	switch t.Params["site"] {
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

	proxyData, err := proxy.GetProxyFromProxyGroup(t.ProxyListID)

	if err != nil {
		return task.ErrorTaskState
	}

	client, err := hclient.NewClient(proxyData.URL, internal.Host)

	if err != nil {
		return task.ErrorTaskState
	}

	t.Client = client

	return GET_SESSION
}
