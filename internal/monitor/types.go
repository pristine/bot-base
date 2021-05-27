package monitor

import (
	"context"
	"github.com/EdwinJ0124/bot-base/third_party/hclient"
	"reflect"
)

type Monitor struct {
	ID          string             `json:"id"`
	Params      map[string]string  `json:"params"`
	Type        string             `json:"type"`
	Input       string             `json:"input"`
	ProxyListID string             `json:"proxyListID"`
	Internal    interface{}        `json:"-"`
	Context     context.Context    `json:"-"`
	Cancel      context.CancelFunc `json:"-"`
	Client      *hclient.Client    `json:"-"`
	Active      bool               `json:"-"`
}

type MonitorType struct {
	firstHandlerState MonitorState
	internalType      reflect.Type
	handlers          MonitorReflectMap
}

type MonitorState string
type MonitorHandlerMap map[MonitorState]interface{}
type MonitorReflectMap map[string]reflect.Value
