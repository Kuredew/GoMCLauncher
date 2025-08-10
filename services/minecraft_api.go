package services

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/utils"
)


func GetVersionManifest() map[string]interface{} {
	log.Print("Getting version Manifest...")
	versionManifestPath := config.AssetVersionManifestFile

	var url string = "https://piston-meta.mojang.com/mc/game/version_manifest_v2.json"
	var data []byte

	data, err := utils.ReadFile(versionManifestPath)

	if err == nil {
		return utils.JsonFormater(data)
	}

	data, err = utils.Response(url)

	if err != nil {
		log.Fatalf("Error while fetching : %s", err)

		log.Print("Retrying...")
		return GetVersionManifest()
	} 

	utils.WriteFile(config.AssetVersionManifestFile, data)

	version_manifest := utils.JsonFormater(data)
	return version_manifest
}

func GetAssetIndex(dataVersion map[string]interface{}) map[string]interface{} {
	log.Print("Getting asset index...")
	var responseJson []byte
	var responseObject map[string]interface{}
	var err error

	responseJson, err = utils.Response(dataVersion["url"].(string))

	if err != nil {
		log.Fatalf("Error while fetching : %s", err)

		log.Print("Retrying...")
		return GetVersionManifest()
	}

	responseObject = utils.JsonFormater(responseJson)["assetIndex"].(map[string]interface{})

	assetFileName := fmt.Sprintf("%s.json", responseObject["id"].(string))
	assetFilePath := filepath.Join(config.AssetIndexDir, assetFileName)

	data, err := utils.ReadFile(assetFilePath)

	if err == nil {
		return utils.JsonFormater(data)
	}

	responseJson, err = utils.Response(responseObject["url"].(string))

	if err != nil {
		log.Fatalf("Error while fetching : %s", err)

		log.Print("Retrying...")
		return GetVersionManifest()
	}

	utils.WriteFile(assetFilePath, responseJson)
	
	return utils.JsonFormater(responseJson)
}