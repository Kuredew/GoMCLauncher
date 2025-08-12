package config

import (
	"path/filepath"
)


var DATA_PATH string = "data"

var InstanceDir string = filepath.Join(DATA_PATH, "instances")

var AssetDir string = filepath.Join(DATA_PATH, "assets")
var AssetIndexDir string = filepath.Join(AssetDir, "indexes")
var AssetObjectDir string = filepath.Join(AssetDir, "objects")
var AssetVersionManifestFile string = filepath.Join(AssetDir, "versionManifest", "version_manifest_v2.json")

