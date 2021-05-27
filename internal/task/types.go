package task

import (
	"context"
	"github.com/EdwinJ0124/bot-base/third_party/hclient"
	"reflect"
)

type Task struct {
	ID             string             `json:"id"`
	Type           string             `json:"type"`
	Site           string             `json:"site"`
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
	internalType      reflect.Type
	handlers          TaskReflectMap
}

type TaskState string
type TaskHandlerMap map[TaskState]interface{}
type TaskReflectMap map[string]reflect.Value
