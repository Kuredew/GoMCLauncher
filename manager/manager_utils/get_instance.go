package managerutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/model"
	"github.com/Kuredew/GoMCLauncher/utils"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

func GetInstance() (model.Instance, error) {
	entries, _ := os.ReadDir(config.InstanceDir)
	var isNewInstance bool = false
	var instancelist []interface{}
	var instance model.Instance

	if len(entries) < 1 {
		return instance, errors.New("no instance")
	}

	for _, entry := range entries {
		instancelist = append(instancelist, entry.Name())
	}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	var choices []string
	selected := 0

	for _, name := range instancelist {
		choices = append(choices, name.(string))
	}

	blue := color.New(color.FgBlue).SprintFunc()
	for {
		utils.ClearScreen()
		fmt.Print("Welcome to GoMCLauncher!\n")
		
		fmt.Print("\n🏡 Choose Instance to start playing minecraft\n\n")
		
		for index, choice := range choices {
			if index == selected {
				fmt.Printf("  %s %s\n", "→", blue(choice))
			} else {
				fmt.Printf("    %s\n", choice)
			}
		}
		fmt.Print("\n\nUse Arrow Keys [↑↓] and enter to select instance\nPress CTRL+N to create New Instance\n\n")

		_, key, _ := keyboard.GetKey()

		if key == keyboard.KeyArrowUp {
			if selected > 0 {
				selected--
			}
		}
		if key == keyboard.KeyArrowDown {
			if selected < len(choices)-1 {
				selected++
			}
		}
		if key == keyboard.KeyEnter {
			break
		}
		if key == keyboard.KeyCtrlN {
			isNewInstance = true

			break
		}
	}

	if isNewInstance {
		CreateNewInstance()

		return GetInstance()
	}

	selectedInstanceName := instancelist[selected].(string)

	// Read config file in selected instance directory
	configJson, _ := os.ReadFile(filepath.Join(config.InstanceDir, selectedInstanceName, "config.json"))
	json.Unmarshal(configJson, &instance)

	return instance, nil
}