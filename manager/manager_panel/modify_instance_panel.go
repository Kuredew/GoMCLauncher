package managerpanel

import (
	"fmt"

	managerutils "github.com/Kuredew/GoMCLauncher/manager/manager_utils"
	"github.com/Kuredew/GoMCLauncher/model"
	"github.com/Kuredew/GoMCLauncher/utils"
)

func ModifyInstancePanel(instance model.Instance) {
	oldInstance := instance

	for {
		options := []string{"Instance Name        " + instance.Name, "Minecraft Version    " + instance.Version, "", "Save", "Cancel"}

		headerString := fmt.Sprintf("Modify %s Instance", instance.Name)

		selected, err := utils.CreatePanel(headerString, options)
		if err != nil {
			return
		}
		
		switch selected {
		case 0:
			managerutils.ChangeInstanceName(&instance)
		case 1:
			managerutils.ChangeVersion(&instance)
		case 3:
			managerutils.SaveModifiedInstance(oldInstance, instance)
			return
		case 4:
			return
		}
	}

}