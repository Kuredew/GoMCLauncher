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

	dataJson, err = utils.Response(url)

	if err != nil {
		log.Fatalf("Error while fetching : %s", err)

		log.Print("Retrying...")
		return GetVersionManifest()
	} 

	utils.WriteFile(config.AssetVersionManifestFile, dataJson)

	json.Unmarshal(dataJson, &data)
	return data
}

func GetDependency(versionInfo map[string]interface{}) ([]interface{}, map[string]interface{}) {
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
		versionDependencyJson, _ = utils.Response(versionInfo["url"].(string))
	}
	json.Unmarshal(versionDependencyJson, &versionDependency)

	// get version asset id from version dependency
	versionAssetId = versionDependency["assetIndex"].(map[string]interface{})["id"].(string)
	log.Printf("AssetIndex : %s", versionAssetId)

	var versionAssetListFilePath = filepath.Join(config.AssetIndexDir, versionAssetId + ".json")

	// get version assets list
	versionAssetListJson, err = utils.ReadFile(versionAssetListFilePath)
	if err != nil {
		versionAssetListJson, _ = utils.Response(versionDependency["assetIndex"].(map[string]interface{})["url"].(string))
	}
	json.Unmarshal(versionAssetListJson, &versionAssetList)
	
	// finally, save to file.
	utils.WriteFile(versionAssetListFilePath, versionAssetListJson)
	utils.WriteFile(versionDependencyFilePath, versionDependencyJson)

	return versionDependency["libraries"].([]interface{}), versionAssetList
}