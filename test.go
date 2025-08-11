package main

import (
	"fmt"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/utils"
)

func mains() {
	hash := "b62ca8ec10d07e6bf5ac8dae0c8c1d2e6a1e3356"
	hashId := string([]rune(hash)[:2])

	url := "https://resources.download.minecraft.net/%s/%s"
	downloadUrl := fmt.Sprintf(url, hashId, hash)

	//fmt.Printf("%s\n", hashId)
	//fmt.Printf("%s\n", downloadUrl)
	utils.Download(filepath.Join("data", "downloadtest", hashId, hash), downloadUrl)

}