package utils

import "os"

func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func FolderExist(folderPath string) bool {
	info, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}