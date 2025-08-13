package managerutils

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/utils"
	//"github.com/Kuredew/GoMCLauncher/utils"
)

func isLibraryDownloaded(libraryPath string) bool {
	_, err := os.Stat(libraryPath)

	return !os.IsNotExist(err)
}

func GetLibraries(dependencyList map[string]interface{}) string {
	libraries := dependencyList["libraries"].([]interface{})

	var classpath []string
	downloadList := make(map[string]map[string]string)
	var libraryNameList []string
	var libraryDownloadPath string
	var libraryDownloadUrl string

	for _, libraryInfo := range libraries {
		var isAllow bool = true

		libraryName := libraryInfo.(map[string]interface{})["name"].(string)

		artifact := libraryInfo.(map[string]interface{})["downloads"].(map[string]interface{})["artifact"].(map[string]interface{})
		//libraryDownloadPath = filepath.Join(config.LibrariesDir, artifact["path"].(string))

		// idk why im do this.
		splittedLibraryName := strings.Split(libraryName, ":")
		var path []string
		firstPath := strings.ReplaceAll(splittedLibraryName[0], ".", "/")
		secondPath := splittedLibraryName[1] + "/" + splittedLibraryName[2]
		thirdPath := splittedLibraryName[1] + "-" + splittedLibraryName[2] + ".jar"
		path = append(path, firstPath)
		path = append(path, secondPath)
		path = append(path, thirdPath)

		libraryDownloadPath = filepath.Join(config.LibrariesDir, strings.Join(path, "/"))
		libraryDownloadUrl = artifact["url"].(string)

		// Handle rules in libraries.
		if rules, ok := libraryInfo.(map[string]interface{})["rules"]; ok {
			for _, rule := range rules.([]interface{}) {
				os, osFound := rule.(map[string]interface{})["os"]
				action, _ := rule.(map[string]interface{})["action"]
				
				if action.(string) == "allow" {
					isAllow = true
				}

				if !osFound {
					continue
				}

				for _, osName := range os.(map[string]interface{}) {
					if action.(string) == "allow" && osName == "windows" {
						isAllow = true

						break
					}

					if action.(string) == "allow" && osName != "windows" {
						isAllow = false
					}

					if action.(string) == "disallow" && osName == "windows" {
						isAllow = false
						log.Printf("Skipping disallow lib %s", libraryName)

						break
					}
				}
			}
		}

		if isAllow {
			// Handle natives libraries for old minecraft.
			if natives, ok := libraryInfo.(map[string]interface{})["natives"]; ok {
				for key, value := range natives.(map[string]interface{}) {
					if key != "windows" {
						continue
					}
					classifiers, ok := libraryInfo.(map[string]interface{})["downloads"].(map[string]interface{})["classifiers"]

					if !ok {
						continue
					}

					nativeDownloadPath := filepath.Join(config.LibrariesDir, classifiers.(map[string]interface{})[value.(string)].(map[string]interface{})["path"].(string))
					nativeDownloadUrl := classifiers.(map[string]interface{})[value.(string)].(map[string]interface{})["url"].(string)

					nativeLibraryName := libraryName + "natives"

					downloadList[nativeLibraryName] = make(map[string]string)
					downloadList[nativeLibraryName][nativeDownloadPath] = nativeDownloadUrl
					libraryNameList = append(libraryNameList, nativeLibraryName)
				}
			}

			downloadList[libraryName] = make(map[string]string)
			downloadList[libraryName][libraryDownloadPath] = libraryDownloadUrl
			libraryNameList = append(libraryNameList, libraryName)
		}
	}
	// Handle Client.
	clientDownloadUrl := dependencyList["downloads"].(map[string]interface{})["client"].(map[string]interface{})["url"].(string)
	clientDownloadPath := filepath.Join(config.DATA_PATH, "versions", dependencyList["id"].(string), dependencyList["id"].(string) + ".jar")

	downloadList["client"] = make(map[string]string)
	downloadList["client"][clientDownloadPath] = clientDownloadUrl
	libraryNameList = append(libraryNameList, "client")

	// Download
	for index, libraryName := range libraryNameList {
		downloadInfo := downloadList[libraryName]
		var path string
		var url string


		for key, value := range downloadInfo {
			path = key
			url = value
		}

		
		classpath = append(classpath, path)
		
		if isLibraryDownloaded(path) {
			continue
		}
		
		utils.Download(path, url)

		log.Printf("[%v/%v] Downloaded", index+1, len(libraryNameList))
	}
	log.Printf("Loaded %v library.", len(libraryNameList))

	return strings.Join(classpath, ";")
}