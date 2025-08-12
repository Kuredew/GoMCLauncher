package managerutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/utils"
	//"github.com/Kuredew/GoMCLauncher/utils"
)

func isLibraryDownloaded(libraryPath string) bool {
	_, err := os.Stat(libraryPath)

	return !os.IsNotExist(err)
}

func GetLibraries(libraries []interface{}) {
	fmt.Print("Downloading Libraries\n")

	var libraryName string
	var libraryDownloadPath string
	var libraryDownloadUrl string

	for _, libraryInfo := range libraries {
		libraryName = libraryInfo.(map[string]interface{})["name"].(string)

		if rules, ok := libraryInfo.(map[string]interface{})["rules"]; ok {
			rule := rules.([]interface{})[0]
			
			if rule.(map[string]interface{})["os"].(map[string]interface{})["name"] != "windows" {
				log.Printf("Skipping %s", libraryName)
				continue
			}
		}

		artifact := libraryInfo.(map[string]interface{})["downloads"].(map[string]interface{})["artifact"].(map[string]interface{})
		libraryDownloadPath = filepath.Join(config.LibrariesDir, artifact["path"].(string))
		libraryDownloadUrl = artifact["url"].(string)

		if isLibraryDownloaded(libraryDownloadPath) {
			continue
		}

		log.Printf("Downloading %s", libraryDownloadUrl)
		utils.Download(libraryDownloadPath, libraryDownloadUrl)
	}

}