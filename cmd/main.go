package main

import (
	"github.com/EdwinJ0124/bot-base/monitors/footsitesMonitor"
	"github.com/EdwinJ0124/bot-base/sites/footsites"
)

func main() {
	footsitesMonitor.Initialize()
	footsites.Initialize()
}
