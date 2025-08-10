package main

import (
	"log"

	"github.com/Kuredew/GoMCLauncher/manager"
	"github.com/Kuredew/GoMCLauncher/utils"
)

var MC_LATEST_VERSION string
var ASSET_INDEX int

func main() {
	instance, err := utils.GetInstanceName()
	if err != nil {
		log.Printf("Error getting Instance : %s", err)

		manager.NewInstance(manager.Instance{Name: "Kureichi Minecraft", Version: "1.21.8", Modloader: "fabric"})
	}

	log.Printf("\nStarting %s", instance)
}