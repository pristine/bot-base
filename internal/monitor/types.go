package monitor

import (
	"context"
	"github.com/EdwinJ0124/bot-base/third_party/hclient"
	"github.com/iancoleman/orderedmap"
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
	handlers          *orderedmap.OrderedMap
}