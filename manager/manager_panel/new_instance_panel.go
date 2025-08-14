package managerpanel

import (
	"encoding/json"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	managerutils "github.com/Kuredew/GoMCLauncher/manager/manager_utils"
	"github.com/Kuredew/GoMCLauncher/model"
	"github.com/Kuredew/GoMCLauncher/utils"
)




func CreateNewInstancePanel() {
	var instance model.Instance

	instance.Modloader = "vanilla"
	instance.AssetIndex = "default"

	var err error
	selected := 0

	Loop:
	for {
		options := []string{"Instance Name        " + instance.Name, "Minecraft Version    " + instance.Version, "ModLoader            " + instance.Modloader, "", "Finish!", "Cancel"}
		selected, err = utils.CreatePanel("✨ Complete this form to create New Instance.", options)
		
		if err != nil {
			return
		}
		switch selected {
			case 0:
				managerutils.ChangeInstanceName(&instance)
			case 1:
				managerutils.ChangeVersion(&instance)
			case 2:
				managerutils.ChangeModloader(&instance)
			case 4:
				if instance.Name == "" || instance.Version == "" || instance.Modloader == "" {
					continue Loop
				}


				instanceJson , _ := json.Marshal(instance)

				utils.WriteFile(filepath.Join(config.InstanceDir, instance.Name, "config.json"), instanceJson)
				break Loop
			case 5: 
				return
		}
	}
}