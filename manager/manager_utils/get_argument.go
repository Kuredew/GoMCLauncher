package managerutils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/model"
)

func GetArg(dependencyInfo map[string]interface{}, classpath string, instance model.Instance) []string {
	arguments := dependencyInfo["arguments"].(map[string]interface{})
	mainClass := dependencyInfo["mainClass"].(string)

	gameArgRaw := arguments["game"].([]interface{})
	JVMArgRaw := arguments["jvm"].([]interface{})

	gameArg := GetGameArg(gameArgRaw, instance)
	JVMArg := GetJavaArg(JVMArgRaw, mainClass, classpath)


	return append(JVMArg, gameArg...)
}

func GetJavaArg(JVMArgRaw []interface{}, mainClass string, classpath string) []string {
	var javaArgument []string
	classpathQuoted := fmt.Sprintf(`"%s"`, classpath)

	os.Setenv("NATIVES_DIRECTORY", config.LibrariesDir)
	os.Setenv("LAUNCHER_NAME", config.LauncherName)
	os.Setenv("LAUNCHER_VERSION", config.LauncherVersion)
	os.Setenv("CLASSPATH", classpathQuoted)
	
	javaArgument = append(javaArgument, "-Xmx4G")
	javaArgument = append(javaArgument, "-Dminecraft.api.env=custom")
	javaArgument = append(javaArgument, "-Dminecraft.api.auth.host=https://invalid.invalid")
	javaArgument = append(javaArgument, "-Dminecraft.api.account.host=https://invalid.invalid")
	javaArgument = append(javaArgument, "-Dminecraft.api.session.host=https://invalid.invalid")
	javaArgument = append(javaArgument, "-Dminecraft.api.services.host=https://invalid.invalid")

	for _, value := range JVMArgRaw {
		if _, ok := value.(map[string]interface{}); ok {
			continue
		}

		finalString := os.ExpandEnv(value.(string))
		javaArgument = append(javaArgument, finalString)
	}

	javaArgument = append(javaArgument, mainClass)

	return javaArgument
}

func GetGameArg(gameArgRaw []interface{}, instance model.Instance) []string {
	var gameArgument []string

	usernameQuoted := fmt.Sprintf(`"%s"`, config.Username)
	versionNameQuoted := fmt.Sprintf(`"%s"`, instance.Version)
	gameDirQuoted := filepath.Join(config.InstanceDir, instance.Name)
	assetDirQuoted := config.AssetDir
	assetIndexNameQuoted := instance.AssetIndex

	os.Setenv("AUTH_PLAYER_NAME", usernameQuoted)
	os.Setenv("VERSION_NAME", versionNameQuoted)
	os.Setenv("GAME_DIRECTORY", gameDirQuoted)
	os.Setenv("ASSETS_ROOT", assetDirQuoted)
	os.Setenv("ASSETS_INDEX_NAME", assetIndexNameQuoted)
	os.Setenv("LAUNCHER_VERSION", config.LauncherVersion)

	os.Setenv("auth_uuid", `""`)
	os.Setenv("auth_access_token", `""`)
	os.Setenv("clientid", `""`)
	os.Setenv("auth_xuid", `""`)
	os.Setenv("user_type", `""`)
	os.Setenv("version_type", `""`)

	for _, value := range gameArgRaw {
		if _, ok := value.(map[string]interface{}); ok {
			continue
		}

		finalString := os.ExpandEnv(value.(string))
		gameArgument = append(gameArgument, finalString)

	}

	return gameArgument
}