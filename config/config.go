package config

import (
	"path/filepath"
)


var DATA_PATH string = "data"

var INSTANCE_PATH_DIR string = filepath.Join(DATA_PATH, "instances")

var ASSET_PATH string = filepath.Join(DATA_PATH, "assets")
var AssetIndexDir string = filepath.Join(ASSET_PATH, "indexes")
var AssetObjectDir string = filepath.Join(ASSET_PATH, "objects")
var AssetVersionManifestFile string = filepath.Join(ASSET_PATH, "versionManifest", "version_manifest_v2.json")
