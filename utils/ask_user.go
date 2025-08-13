package utils

import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"

	"github.com/fatih/color"
)

func AskUserInput() string {
	blue := color.New(color.FgBlue).SprintFunc()

	fmt.Printf("%s", blue("\n→ "))

	var userInput string
	fmt.Scan(&userInput)
	//userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	//userInput = strings.ReplaceAll(userInput, "\n", "")

	return userInput
}
