package manager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/services"
)

type Instance struct {
	Name      string
	Version   string
	Modloader string
}

func NewInstance(instance Instance) bool {
	newInstancePath := filepath.Join(config.INSTANCE_PATH_DIR, instance.Name)

	os.MkdirAll(newInstancePath, 0755)

	// search version in version manifest
	version_manifest := services.GetVersionManifest()

	version_list := version_manifest["versions"].([]interface{})

	var data map[string]interface{}
	for _, value := range version_list {
		id := value.(map[string]interface{})["id"].(string)

		if id == instance.Version {
			data = services.GetAssetIndex(value.(map[string]interface{}))
		}
	}
	fmt.Print(data)

	return true
}

func StartInstance(instance Instance) bool {
	fmt.Print(instance.Name)

	return true
}