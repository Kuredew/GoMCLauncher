package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/manager"
	"github.com/Kuredew/GoMCLauncher/model"
	"github.com/Kuredew/GoMCLauncher/panel"
	"github.com/Kuredew/GoMCLauncher/utils"
	"github.com/fatih/color"
)

func getConfig() model.Config {
	var configModel model.Config

	configJson, err := utils.ReadFile(config.ConfigFile)
	if err == nil {
		json.Unmarshal(configJson, &configModel)
	} else {
		panel.SetOptionsPanel(&configModel)
	}

	return configModel
}

func main() {
	panelConfig := getConfig()

	fmt.Print("\033[H\033[2J")
	options := []string{"Select Instances", "Options", "Quit"}

	blue:= color.New(color.FgBlue).SprintFunc()
	Loop:
	for {
		headerPanel := fmt.Sprintf("✨ Hello, %s! Welcome to GoMCLauncher\nSelect Options :", blue(panelConfig.PlayerName))

		userSelected, _ := utils.CreatePanel(headerPanel, options)

		switch userSelected {
		case 0:
			err := manager.Initialize(panelConfig)
			if errors.Is(err, manager.ErrBack) {
				continue Loop
			}

			break Loop
		case 1:
			panel.SetOptionsPanel(&panelConfig)
		case 2:
			os.Exit(0)
		}
	}
}