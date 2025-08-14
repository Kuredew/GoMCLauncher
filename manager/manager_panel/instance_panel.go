package managerpanel

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	managerutils "github.com/Kuredew/GoMCLauncher/manager/manager_utils"
	"github.com/Kuredew/GoMCLauncher/model"
	"github.com/Kuredew/GoMCLauncher/utils"
)

var ErrInstancePanelBack = errors.New("back")
var ErrInstancePanelNewInstance = errors.New("newInstance")
var ErrInstancePanelNoInstance = errors.New("noInstance")

func InstancePanel() (model.Instance, error) {
	instanceNameList := managerutils.GetInstances()
	var instance model.Instance

	if len(instanceNameList) < 1 {
		return instance, ErrInstancePanelNoInstance
	}

	instanceNameList = append(instanceNameList, "", "Create New Instance", "Back")
	
	userSelected, err := utils.CreatePanel("🏡 Choose Instances", instanceNameList)

	if err != nil {
		return InstancePanel()
	}

	if userSelected == len(instanceNameList)-2 {
		return instance, ErrInstancePanelNewInstance
	} else if userSelected == len(instanceNameList)-1 {
		return instance, ErrInstancePanelBack
	}

	selectedInstanceName := instanceNameList[userSelected]

	// Read config file in selected instance directory
	configJson, _ := os.ReadFile(filepath.Join(config.InstanceDir, selectedInstanceName, "config.json"))
	json.Unmarshal(configJson, &instance)

	return instance, nil
}