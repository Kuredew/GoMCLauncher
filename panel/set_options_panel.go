package panel

import (
	"encoding/json"
	"fmt"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/model"
	"github.com/Kuredew/GoMCLauncher/utils"
)
func saveOptions(configModel *model.Config) {
	configJson, _ := json.Marshal(*configModel)

	utils.WriteFile(config.ConfigFile, configJson)
}

func setPlayerName(configModelAddress *model.Config) {
	fmt.Print("Write the username you want")
	userInput := utils.AskUserInput()

	configModel := *configModelAddress
	configModel.PlayerName = userInput

	*configModelAddress = configModel
}

func SetOptionsPanel(configModelAddress *model.Config) {
	OptionLoop:
	for {
		configModel := *configModelAddress

		options := []string{"Player Name        " + configModel.PlayerName, "", "Save", "Cancel"}
		userSelected, _ := utils.CreatePanel("These are the settings you need to fill in", options)

		switch userSelected {
		case 0:
			setPlayerName(configModelAddress)
		case 2:
			if configModel.PlayerName == "" {
				continue OptionLoop
			}

			saveOptions(configModelAddress)
			return
		case 3:
			if configModel.PlayerName == "" {
				continue OptionLoop
			}
			
			return
		}
	}
}