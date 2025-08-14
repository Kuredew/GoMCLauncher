package managerutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/model"
	"github.com/Kuredew/GoMCLauncher/services"
	"github.com/Kuredew/GoMCLauncher/utils"
)

// For save and delete instance
func SaveModifiedInstance(oldInstance model.Instance, newInstance model.Instance) {
	oldInstanceConfigPath := filepath.Join(config.InstanceDir, oldInstance.Name, "config.json")
	newInstanceConfigPath := filepath.Join(config.InstanceDir, newInstance.Name, "config.json")

	// Change config json first
	newConfigJson, _ := json.Marshal(newInstance)
	utils.WriteFile(oldInstanceConfigPath, newConfigJson)

	// Rename directory
	os.Rename(filepath.Dir(oldInstanceConfigPath), filepath.Dir(newInstanceConfigPath))
}
func DeleteInstances(instance model.Instance) error {
	fmt.Print(`Are you sure?! all data in this instance will be lost! (type "yes" to delete this instance)`)
	userType := utils.AskUserInput()

	if userType != "yes" {
		return errors.New("cancel")
	}

	instanceName := instance.Name
	instancePath := filepath.Join(config.InstanceDir, instanceName)

	return os.RemoveAll(instancePath)
}


// For change instance value.
func ChangeInstanceName(instanceAddress *model.Instance) {
	fmt.Print(`Give a name to your instance like "my-instance" (space is not allowed.)`)
	instanceAddress.Name = utils.AskUserInput()
}
func ChangeVersion(instanceAddress *model.Instance) {
	fmt.Print(`Please type the version of Minecraft you want to play, such as "1.21.8"`)
	userInput := utils.AskUserInput()

	versionManifest := services.GetVersionManifest()
	versionList := versionManifest["versions"].([]interface{})

	isValid := false
	for _, version := range versionList {
		id := version.(map[string]interface{})["id"].(string)

		if id == userInput {
			fmt.Print("Closed\n\n")
			isValid = true
		}
	}

	if !isValid {
		return
	}
	instanceAddress.Version = userInput
}
func ChangeModloader(instanceAddress *model.Instance) {
	options := []string{"vanilla (default)", "forge", "neoforge", "fabric"}

	userSelected, err := utils.CreatePanel("\nChoose Modloader\n\n", options)
	if err != nil {
		return
	}

	if options[userSelected] == "vanilla (default)" {
		options[userSelected] = "vanilla"
	}

	instanceAddress.Modloader = options[userSelected]
}



func RenameInstanceName(instance model.Instance, newInstanceName string) error {
	instanceName := instance.Name
	instancePath := filepath.Join(config.InstanceDir, instanceName)

	newInstancePath := filepath.Join(config.InstanceDir, newInstanceName)
	newInstanceConfigPath := filepath.Join(newInstancePath, "config.json")
	
	err := os.Rename(instancePath, newInstancePath)
	if err != nil {
		return err
	}
	
	instance.Name = newInstanceName
	configJson, _ := json.Marshal(instance)

	utils.WriteFile(newInstanceConfigPath, configJson)

	return nil
}

func ChangeInstanceMinecraftVersion(instance model.Instance, newVersion string) error {
	instanceName := instance.Name
	instance.Version = newVersion

	instancePath := filepath.Join(config.InstanceDir, instanceName)
	InstanceConfigPath := filepath.Join(instancePath, "config.json")

	configJson, err := json.Marshal(instance)
	if err != nil {
		return err
	}

	utils.WriteFile(InstanceConfigPath, configJson)
	return nil
}