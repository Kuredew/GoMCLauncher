package manager

import (
	"os"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	managerutils "github.com/Kuredew/GoMCLauncher/manager/manager_utils"
	"github.com/Kuredew/GoMCLauncher/model"
	"github.com/Kuredew/GoMCLauncher/services"
)

func getDependency(instance model.Instance) {
	var libraries []interface{}
	var assetList map[string]interface{}

	versionManifest := services.GetVersionManifest()
	versionList := versionManifest["versions"].([]interface{})
	
	// search version in version manifest
	for _, value := range versionList {
		id := value.(map[string]interface{})["id"].(string)

		if id == instance.Version {
			libraries, assetList = services.GetDependency(value.(map[string]interface{}))

			managerutils.GetLibraries(libraries)
			managerutils.GetAsset(assetList)
		}
	}
}

func CreateNewInstance(instance model.Instance) bool {
	newInstancePath := filepath.Join(config.InstanceDir, instance.Name)
	os.MkdirAll(newInstancePath, 0755)

	getDependency(instance)

	return true
}

func StartInstance(instance model.Instance) bool {
	instancePath := filepath.Join(config.InstanceDir, instance.Name)
	os.MkdirAll(instancePath, 0755)

	getDependency(instance)

	return true
}