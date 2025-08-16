package utils

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

func CreatePanel(headerString string, choices []string) (int, error) {
	keyboard.Open()
	defer keyboard.Close()

	selected := 0

	blue := color.New(color.BgBlue).SprintFunc()
	for {
		ClearScreen()

		fmt.Print("\n" + headerString + "\n\n")

		for i, choice := range choices {
			if i == selected {
				fmt.Printf("  %s %s\n", "→", blue(choice))
			} else {
				fmt.Printf("    %s\n", choice)
			}
		}
		fmt.Print("\n\nUse Arrow Keys [↑↓] and enter to select value\n\n")

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
			return selected, nil
		}
	}

}