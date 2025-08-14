package main

import (
	"fmt"
	"os"

	"github.com/Kuredew/GoMCLauncher/manager"
	"github.com/Kuredew/GoMCLauncher/utils"
)

func main() {
	fmt.Print("\033[H\033[2J")
	options := []string{"Select Instances", "Quit"}

	for {
		userSelected, _ := utils.CreatePanel("\n✨ Welcome to TuiMC\nSelect Options :\n\n", options)

		switch userSelected {
		case 0:
			manager.Initialize()
		case 1:
			os.Exit(0)
		}
	}
}