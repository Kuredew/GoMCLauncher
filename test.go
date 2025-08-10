package main

import "github.com/Kuredew/GoMCLauncher/manager"

func mains() {
	instance := manager.Instance{Name: "KureichiMinecraft", Version: "1.21.8", Modloader: "Fabric"}
	
	manager.StartInstance(instance)
}