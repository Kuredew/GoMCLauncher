package managerutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/model"
	"github.com/Kuredew/GoMCLauncher/services"
	"github.com/Kuredew/GoMCLauncher/utils"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

var instance model.Instance


func changeInstanceName() {
	//utils.ClearScreen()

	fmt.Print(`Give a name to your instance like "my-instance" (space is not allowed.)`)
	instance.Name = utils.AskUserInput()
}

func changeVersion() {
	//utils.ClearScreen()

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
	instance.Version = userInput
}

func changeModloader() {
	//utils.ClearScreen()

	fmt.Print(`This is optional, available options are "fabric", "forge" and "neoforge" (type "vanilla" to default)`)
	userInput := utils.AskUserInput()

	if userInput != "vanille" && userInput != "fabric" && userInput != "forge" && userInput != "neoforge" {
		instance.Modloader = "vanilla"
		return
	}
	instance.Modloader = userInput
}


func CreateNewInstance() error {
	instance.Modloader = "vanilla"
	instance.AssetIndex = "default"

	selected := 0

	keyboard.Open()
	defer keyboard.Close()

	blue := color.New(color.FgBlue).SprintFunc()
	for {
		options := []string{"Instance Name        " + instance.Name, "Minecraft Version    " + instance.Version, "ModLoader            " + instance.Modloader, "", "Finish!"}

		utils.ClearScreen()

		fmt.Print("✨ Complete this form to create New Instance.\n\n")

		for i, option := range options {
			if i == selected {
				fmt.Printf("  %s %s\n", "→", blue(option))
			} else {
				fmt.Printf("    %s\n", option)
			}
		}
		fmt.Print("\n\nUse Arrow Keys [↑↓] and enter to change value\nPress ESC to Back\n\n")

		_, key, _ := keyboard.GetKey()

		if key == keyboard.KeyArrowUp {
			if selected > 0 {
				selected--
			}
		}
		if key == keyboard.KeyArrowDown {
			if selected < len(options)-1 {
				selected++
			}
		}
		if key == keyboard.KeyEnter {
			if selected == 0 {
				changeInstanceName()
			}
			if selected == 1 {
				changeVersion()
			}
			if selected == 2 {
				changeModloader()
			}
			if selected == 4 && instance.Name != "" && instance.Version != "" && instance.Modloader != "" {
				instanceJson , _ := json.Marshal(instance)

				utils.WriteFile(filepath.Join(config.InstanceDir, instance.Name, "config.json"), instanceJson)
				break
			}
		}
		if key == keyboard.KeyEsc {
			return errors.New("back")
		}
	}
	return nil
}