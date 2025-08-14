package managerutils

import (
	"os"
	"slices"

	"github.com/Kuredew/GoMCLauncher/config"
)

func GetInstances() []string {
	var instanceNameSlice []string

	instances, _ := os.ReadDir(config.InstanceDir)
	//instancesInfo := make(map[string]string)

	for _, instanceName := range instances {		
		instanceNameSlice = append(instanceNameSlice, instanceName.Name())
	}
	slices.Sort(instanceNameSlice)

	/*
	for _, instanceName := range instanceNameSlice {
		instancesInfo[instanceName] = filepath.Join(config.InstanceDir, instanceName)
	}*/

	return instanceNameSlice
}