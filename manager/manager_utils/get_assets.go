package managerutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/utils"
)

func isAssetDownloaded(assetPath string) bool {
	_, err := os.Stat(assetPath)

	return !os.IsNotExist(err)
}

func GetAsset(assetList map[string]interface{}) {
	var downloadUrl string = "https://resources.download.minecraft.net/%s/%s"
	var objectPath string

	object := assetList["objects"].(map[string]interface{})

	var sortedObjectKey []string

	for key := range object {
		sortedObjectKey = append(sortedObjectKey, key)
	}

	sort.Strings(sortedObjectKey)

	var assetIndex int = 0
	for _, key := range sortedObjectKey {
		assetIndex += 1

		value := object[key]

		objectHash := value.(map[string]interface{})["hash"].(string)
		objectHashId := string([]rune(objectHash)[:2])

		objectPath = filepath.Join(config.AssetObjectDir, objectHashId, objectHash)
		
		objectDownloadUrl := fmt.Sprintf(downloadUrl, objectHashId, objectHash)
		if isAssetDownloaded(objectPath) {
			continue
		}

		utils.Download(objectPath, objectDownloadUrl)

		log.Printf("[%v/%v] Downloaded", assetIndex, len(sortedObjectKey))
	}
	log.Printf("Loaded %v assets.", len(sortedObjectKey))
}