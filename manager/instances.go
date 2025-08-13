package manager

import (
	//"bufio"

	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	managerutils "github.com/Kuredew/GoMCLauncher/manager/manager_utils"
	"github.com/Kuredew/GoMCLauncher/model"
	"github.com/Kuredew/GoMCLauncher/services"
)

var Argument []string
var gameDir string

func getDependency(instance model.Instance) {
	var dependencyInfo map[string]interface{}
	var assetList map[string]interface{}
	var assetIndex string
	var classpath string

	versionManifest := services.GetVersionManifest()
	versionList := versionManifest["versions"].([]interface{})
	
	// search version in version manifest
	for _, value := range versionList {
		id := value.(map[string]interface{})["id"].(string)

		if id == instance.Version {
			dependencyInfo, assetList, assetIndex = services.GetDependency(value.(map[string]interface{}))
			instance.AssetIndex = assetIndex
			
			classpath = managerutils.GetLibraries(dependencyInfo)
			Argument = managerutils.GetArg(dependencyInfo, classpath, instance)

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
	gameDir = filepath.Join(config.InstanceDir, instance.Name)
	os.MkdirAll(gameDir, 0755)

	getDependency(instance)

	log.Printf("Launching %s...\n\n", instance.Name)
	cmd := exec.Command(config.JavaPath, Argument...)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	
	if err := cmd.Start(); err != nil {
		log.Printf("Error calling java client : %s", err)
	}

	go func() {
		io.Copy(os.Stdout, stdout)
	}()

	go func() {
		io.Copy(os.Stderr, stderr)
	}()

	cmd.Wait()
	
	return true
}