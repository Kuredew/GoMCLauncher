package main

import (
	"log"

	"github.com/Kuredew/GoMCLauncher/manager"
	managerutils "github.com/Kuredew/GoMCLauncher/manager/manager_utils"
	"github.com/Kuredew/GoMCLauncher/model"
)

var MC_LATEST_VERSION string
var ASSET_INDEX int

func main() {
	instance, err := managerutils.GetInstance()
	if err != nil {
		log.Printf("Error getting Instance : %s", err)

		manager.CreateNewInstance(model.Instance{Name: "Kureichi Minecraft", Version: "1.21.8", Modloader: "fabric"})
	}

	log.Printf("\nStarting %s", instance.Name)

	manager.StartInstance(instance)
}