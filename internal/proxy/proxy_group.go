package proxy

import (
	"errors"
	"github.com/iancoleman/orderedmap"
	"github.com/lithammer/shortuuid"
	"sync"
)

type ProxyGroup struct {
	ID      string                 `json:"id"`
	Name    string                 `json:"name"`
	Proxies *orderedmap.OrderedMap `json:"proxies"`
}

var (
	proxyGroupMutex = sync.RWMutex{}

	ProxyGroupEmptyErr        = errors.New("proxy group does not contain any proxies")
	ProxyGroupDoesNotExistErr = errors.New("proxy group does not exist")
	proxyGroups               = make(map[string]*ProxyGroup)
)

// DoesProxyGroupExist checks if a proxy group exists
func DoesProxyGroupExist(id string) bool {
	proxyGroupMutex.RLock()
	defer proxyGroupMutex.RUnlock()
	_, ok := proxyGroups[id]
	return ok
}

// CreateProxyGroup creates a new proxy group
func CreateProxyGroup(name string) string {
	proxyGroupMutex.Lock()
	defer proxyGroupMutex.Unlock()
	id := shortuuid.New()

	proxyGroups[id] = &ProxyGroup{
		ID:      id,
		Name:    name,
		Proxies: orderedmap.New(),
	}

	return id
}

// RemoveProxyGroup removes a proxy group
func RemoveProxyGroup(id string) error {
	if !DoesProxyGroupExist(id) {
		return ProxyGroupDoesNotExistErr
	}

	proxyGroupMutex.Lock()
	defer proxyGroupMutex.Unlock()

	delete(proxyGroups, id)

	return nil
}

// GetProxyFromProxyGroup gets a proxy from a group
func GetProxyFromProxyGroup(id string) (*Proxy, error) {
	if !DoesProxyGroupExist(id) {
		return &Proxy{}, ProxyGroupDoesNotExistErr
	}

	proxyGroupMutex.Lock()
	defer proxyGroupMutex.Unlock()

	proxyGroup := proxyGroups[id]

	proxyIds := proxyGroup.Proxies.Keys()

	if len(proxyIds) == 0 {
		return &Proxy{}, ProxyGroupEmptyErr
	}

	firstProxyId := proxyIds[0]

	proxy, err := GetProxy(firstProxyId)

	if err != nil {
		return &Proxy{}, err
	}

	// remove proxy from list
	proxyGroup.Proxies.Delete(firstProxyId)

	// add proxy to the back of the list
	proxyGroup.Proxies.Set(firstProxyId, true)

	return proxy, nil
}

// GetProxyGroup gets a proxy group
func GetProxyGroup(id string) (*ProxyGroup, error) {
	if !DoesProxyGroupExist(id) {
		return &ProxyGroup{}, ProxyGroupDoesNotExistErr
	}

	proxyGroupMutex.RLock()
	defer proxyGroupMutex.RUnlock()

	proxyGroup := proxyGroups[id]

	return proxyGroup, nil
}

// GetAllProxyGroupIds gets all proxy group ids
func GetAllProxyGroupIds() []string {
	ids := make([]string, 0)

	for id := range proxyGroups {
		ids = append(ids, id)
	}

	return ids
}
