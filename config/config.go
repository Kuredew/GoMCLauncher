package config

import (
	"path/filepath"
)

var LauncherName string = "GoMCLauncher"
var LauncherVersion string = "1.0.0"

var Username string = "Kureichi"

// Java Options
var JavaPath string = filepath.Clean("D:/Download/INSTALLER PROGRAMS/JAVA/jdk-21.0.6/bin/javaw.exe")


// Data Directory
var DATA_PATH string = "data"

var InstanceDir string = filepath.Join(DATA_PATH, "instances")

var AssetDir string = filepath.Join(DATA_PATH, "assets")
var AssetIndexDir string = filepath.Join(AssetDir, "indexes")
var AssetObjectDir string = filepath.Join(AssetDir, "objects")
var AssetVersionManifestFile string = filepath.Join(AssetDir, "versionManifest", "version_manifest_v2.json")

var LibrariesDir string = filepath.Join(DATA_PATH, "libraries")