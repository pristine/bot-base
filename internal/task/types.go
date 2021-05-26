package task

import (
	"context"
	"github.com/EdwinJ0124/bot-base/third_party/hclient"
	"github.com/iancoleman/orderedmap"
)

type Task struct {
	ID             string             `json:"id"`
	Params         map[string]string  `json:"params"`
	Type           string             `json:"type"`
	ProfileGroupID string             `json:"profileGroupID"`
	ProxyListID    string             `json:"proxyListID"`
	Context        context.Context    `json:"-"`
	Cancel         context.CancelFunc `json:"-"`
	Internal       interface{}        `json:"-"`
	Client         *hclient.Client    `json:"-"`
	Active         bool               `json:"-"`
	MonitorData    chan interface{}   `json:"-"`
}


type TaskGroup struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Monitors map[string]bool `json:"monitorId"`
	Tasks    map[string]bool `json:"tasks"`
}

type TaskType struct {
	firstHandlerState TaskState
	handlers          *orderedmap.OrderedMap
}
