package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Kuredew/GoMCLauncher/config"
)

func GetInstanceName() (string, error) {
	entries, _ := os.ReadDir(config.INSTANCE_PATH_DIR)
	var instancelist []interface{}

	if len(entries) < 1 {
		return "", errors.New("no instance")
	}

	for _, entry := range entries {
		instancelist = append(instancelist, entry.Name())
	}

	//return instancelist, nil

	fmt.Print("\n\nChoose Instance to start playing minecraft\n\n")
	for index, name := range instancelist {
		index := fmt.Sprint(index + 1)

		fmt.Printf("%s. %s\n", index, name)
	}
	fmt.Print("\n> ")
	var userInput string
	fmt.Scan(&userInput)

	userInputInt, _ := strconv.Atoi(userInput)

	return instancelist[userInputInt - 1].(string), nil
}