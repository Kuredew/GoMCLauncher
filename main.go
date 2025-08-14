package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Kuredew/GoMCLauncher/manager"
	"github.com/Kuredew/GoMCLauncher/utils"
)

func main() {
	fmt.Print("\033[H\033[2J")
	options := []string{"Select Instances", "Quit"}

Loop:
	for {
		userSelected, _ := utils.CreatePanel("\n✨ Welcome to TuiMC\nSelect Options :\n\n", options)

		switch userSelected {
		case 0:
			err := manager.Initialize()
			if errors.Is(err, manager.ErrBack) {
				continue Loop
			}

			break Loop
		case 1:
			os.Exit(0)
		}
	}
}