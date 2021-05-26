package proxy

import (
	"errors"
	"fmt"
	"github.com/lithammer/shortuuid"
	"strings"
	"sync"
)

type Proxy struct {
	URL string `json:"url"`
}

var (
	proxyMutex = sync.RWMutex{}

	ProxyDoesNotExistErr = errors.New("proxy does not exist")
	proxies              = make(map[string]*Proxy)
)

// DoesProxyExist checks if a proxy exists
func DoesProxyExist(id string) bool {
	proxyMutex.RLock()
	defer proxyMutex.RUnlock()
	_, ok := proxies[id]
	return ok
}

func proxyToProxyUrl(proxy string) string {
	proxySplit := strings.Split(proxy, ":")

	if len(proxySplit) == 2 {
		return fmt.Sprintf("http://%s:%s", proxySplit[0], proxySplit[1])
	} else if len(proxySplit) == 4 {
		return fmt.Sprintf("http://%s:%s@%s:%s", proxySplit[2], proxySplit[3], proxySplit[0], proxySplit[1])
	}

	return fmt.Sprintf("http://%s", proxy)
}

// CreateProxy creates a proxy
func CreateProxy(proxy string) string {
	proxyMutex.Lock()
	defer proxyMutex.Unlock()

	id := shortuuid.New()

	proxies[id] = &Proxy{
		URL: proxyToProxyUrl(proxy),
	}

	return id
}

// RemoveProxy removes a proxy
func RemoveProxy(id string) error {
	if !DoesProxyExist(id) {
		return ProxyDoesNotExistErr
	}

	proxyMutex.Lock()
	defer proxyMutex.Unlock()

	delete(proxies, id)
	return nil
}

func GetProxy(id string) (*Proxy, error) {
	if !DoesProxyExist(id) {
		return &Proxy{}, ProxyDoesNotExistErr
	}

	proxyMutex.RLock()
	defer proxyMutex.RUnlock()

	return proxies[id], nil
}

// AssignProxyToProxyGroup assigns a proxy to a proxy group
func AssignProxyToProxyGroup(proxyId, proxyGroupId string) error {
	if !DoesProxyExist(proxyId) {
		return ProxyDoesNotExistErr
	}

	if !DoesProxyGroupExist(proxyGroupId) {
		return ProxyGroupDoesNotExistErr
	}

	proxyGroup := proxyGroups[proxyGroupId]

	proxyGroup.Proxies.Set(proxyId, true)

	return nil
}
