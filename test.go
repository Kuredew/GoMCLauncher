package main

import (
	"github.com/Kuredew/GoMCLauncher/manager"
	"github.com/Kuredew/GoMCLauncher/model"
)

func mains() {
	instance := model.Instance{Name: "KureichiMinecraft", Version: "1.21.8", Modloader: "Fabric"}
	
	manager.StartInstance(instance)
}