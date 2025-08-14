package main

import (
	"fmt"
	"log"

	"github.com/Kuredew/GoMCLauncher/manager"
	managerpanel "github.com/Kuredew/GoMCLauncher/manager/manager_panel"
	"github.com/Kuredew/GoMCLauncher/model"
)

var MC_LATEST_VERSION string
var ASSET_INDEX int

func main() {
	fmt.Print("\033[H\033[2J")

	instance, err := managerpanel.GetInstancePanel()
	if err != nil {
		log.Printf("Error getting Instance : %s", err)

		manager.CreateNewInstance(model.Instance{Name: "Kureichi Minecraft", Version: "1.21.8", Modloader: "fabric"})
	}

	log.Printf("Initializing %s", instance.Name)

	manager.StartInstance(instance)
}