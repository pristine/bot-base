package main

import (
	"github.com/EdwinJ0124/footsites/monitors/footsitesMonitor"
	"github.com/EdwinJ0124/footsites/sites/footsites"
)

func main() {
	footsitesMonitor.Initialize()
	footsites.Initialize()
}