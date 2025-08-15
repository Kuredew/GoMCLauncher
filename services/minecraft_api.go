package services

import (
	"encoding/json"
	"log"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/utils"
)


func GetVersionManifest() map[string]interface{} {
	log.Print("Getting version Manifest...")
	versionManifestPath := config.AssetVersionManifestFile

	var url string = "https://piston-meta.mojang.com/mc/game/version_manifest_v2.json"
	var data map[string]interface{}

	dataJson, err := utils.ReadFile(versionManifestPath)

	if err == nil {
		json.Unmarshal(dataJson, &data)
		return data
	}
	
	dataJson = utils.Response(url)

	utils.WriteFile(config.AssetVersionManifestFile, dataJson)

	json.Unmarshal(dataJson, &data)
	return data
}

func GetDependency(versionInfo map[string]interface{}) (map[string]interface{}, map[string]interface{}, string) {
	log.Print("Checking Dependency...")
	
	var versionId string
	var versionAssetId string
	var versionDependency map[string]interface{}
	var versionAssetList map[string]interface{}

	var versionAssetListJson []byte
	var versionDependencyJson []byte

	var err error

	// get asset id
	versionId = versionInfo["id"].(string)
	log.Printf("Version ID : %s", versionId)

	var versionDependencyFilePath = filepath.Join(config.AssetIndexDir, versionId, versionId + ".json")

	// get version dependency
	versionDependencyJson, err = utils.ReadFile(versionDependencyFilePath)
	if err != nil {
		versionDependencyJson = utils.Response(versionInfo["url"].(string))
		utils.WriteFile(versionDependencyFilePath, versionDependencyJson)
	}
	json.Unmarshal(versionDependencyJson, &versionDependency)

	// get version asset id from version dependency
	versionAssetId = versionDependency["assetIndex"].(map[string]interface{})["id"].(string)
	log.Printf("AssetIndex : %s", versionAssetId)

	var versionAssetListFilePath = filepath.Join(config.AssetIndexDir, versionAssetId + ".json")

	// get version assets list
	versionAssetListJson, err = utils.ReadFile(versionAssetListFilePath)
	if err != nil {
		versionAssetListJson = utils.Response(versionDependency["assetIndex"].(map[string]interface{})["url"].(string))
		utils.WriteFile(versionAssetListFilePath, versionAssetListJson)
	}
	json.Unmarshal(versionAssetListJson, &versionAssetList)

	return versionDependency, versionAssetList, versionAssetId
}