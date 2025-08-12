package managerutils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"errors"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/model"
)

func GetInstance() (model.Instance, error) {
	entries, _ := os.ReadDir(config.InstanceDir)
	var instancelist []interface{}
	var instance model.Instance

	if len(entries) < 1 {
		return instance, errors.New("no instance")
	}

	for _, entry := range entries {
		instancelist = append(instancelist, entry.Name())
	}

	fmt.Print("\n\nChoose Instance to start playing minecraft\n\n")
	for index, name := range instancelist {
		index := fmt.Sprint(index + 1)

		fmt.Printf("%s. %s\n", index, name)
	}
	fmt.Print("\n> ")
	var userInput string
	fmt.Scan(&userInput)

	userInputInt, _ := strconv.Atoi(userInput)

	selectedInstanceName := instancelist[userInputInt - 1].(string)

	// Read config file in selected instance directory
	configJson, _ := os.ReadFile(filepath.Join(config.InstanceDir, selectedInstanceName, "config.json"))

	json.Unmarshal(configJson, &instance)

	return instance, nil
}