package config

import (
	"path/filepath"
)
var ConfigFile string = "config/config.json"

var LauncherName string = "GoMCLauncher"
var LauncherVersion string = "1.0.0"

var Username string = "Kureichi"

// Data Directory
var DATA_PATH string = "data"

// Java Options
var RuntimeDir string = filepath.Join(DATA_PATH, "runtimes")

var JavaWinDownloadUrl = "https://download.oracle.com/java/21/latest/jdk-21_windows-x64_bin.zip"
var JavaLinuxDownloadUrl = "https://download.oracle.com/java/21/latest/jdk-21_linux-x64_bin.deb"
var JavaMacDownloadUrl = "https://download.oracle.com/java/21/latest/jdk-21_macos-x64_bin.tar.gz"

var JavaWinRuntimeArchive string = filepath.Join(RuntimeDir, "java21.zip")
var JavaLinuxRuntimeArchive string = filepath.Join(RuntimeDir, "java21.tar.gz")
var JavaMacRuntimeArchive string = filepath.Join(RuntimeDir, "java21.tar.gz")

var JavaRuntimeDir string = filepath.Join(RuntimeDir, "java-runtime")

var InstanceDir string = filepath.Join(DATA_PATH, "instances")

var AssetDir string = filepath.Join(DATA_PATH, "assets")
var AssetIndexDir string = filepath.Join(AssetDir, "indexes")
var AssetObjectDir string = filepath.Join(AssetDir, "objects")
var AssetVersionManifestFile string = filepath.Join(AssetDir, "versionManifest", "version_manifest_v2.json")

var LibrariesDir string = filepath.Join(DATA_PATH, "libraries")

var NativeLibrariesDir string = filepath.Join(DATA_PATH, "libraries", "natives")