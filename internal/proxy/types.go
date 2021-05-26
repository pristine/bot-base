package proxy

import "github.com/iancoleman/orderedmap"

type Proxy struct {
	URL string `json:"url"`
}

type ProxyGroup struct {
	ID      string                 `json:"id"`
	Name    string                 `json:"name"`
	Proxies *orderedmap.OrderedMap `json:"proxies"`
}
