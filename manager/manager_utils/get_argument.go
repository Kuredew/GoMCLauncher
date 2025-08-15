package managerutils

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"strings"

	"github.com/Kuredew/GoMCLauncher/config"
	"github.com/Kuredew/GoMCLauncher/model"
)

func GetArg(dependencyInfo map[string]interface{}, classpath string, instance model.Instance, configModel model.Config) []string {
	var arguments = make(map[string]interface{})

	args, ok := dependencyInfo["arguments"].(map[string]interface{})
	if !ok {
		//var gameArguments []string
		var gameArgRaw []interface{}
		var JVMArgRaw []interface{}

		minecraftArguments := strings.Split(dependencyInfo["minecraftArguments"].(string), " ")
		
		for _, minecraftArgument := range minecraftArguments {
			gameArgRaw = append(gameArgRaw, minecraftArgument)
		}

		JVMArgRaw = append(JVMArgRaw, "-Djava.library.path=${natives_directory}", "-Djna.tmpdir=${natives_directory}", "-Dorg.lwjgl.system.SharedLibraryExtractPath=${natives_directory}", "-Dio.netty.native.workdir=${natives_directory}", "-Dminecraft.launcher.brand=${launcher_name}", "-Dminecraft.launcher.version=${launcher_version}", "-cp", "${classpath}")

		arguments["game"] = gameArgRaw
		arguments["jvm"] = JVMArgRaw
	} else {
		maps.Copy(arguments, args)
	}

	mainClass := dependencyInfo["mainClass"].(string)

	gameArgRaw := arguments["game"].([]interface{})
	JVMArgRaw := arguments["jvm"].([]interface{})

	gameArg := GetGameArg(gameArgRaw, instance, configModel)
	JVMArg := GetJavaArg(JVMArgRaw, mainClass, classpath)


	return append(JVMArg, gameArg...)
}

func GetJavaArg(JVMArgRaw []interface{}, mainClass string, classpath string) []string {
	var javaArgument []string
	classpathQuoted := fmt.Sprintf(`"%s"`, classpath)

	os.Setenv("NATIVES_DIRECTORY", config.NativeLibrariesDir)
	os.Setenv("LAUNCHER_NAME", config.LauncherName)
	os.Setenv("LAUNCHER_VERSION", config.LauncherVersion)
	os.Setenv("CLASSPATH", classpathQuoted)
	
	javaArgument = append(javaArgument, "-Xmx4G")
	javaArgument = append(javaArgument, "-Dminecraft.api.env=custom")
	javaArgument = append(javaArgument, "-Dminecraft.api.auth.host=https://invalid.invalid")
	javaArgument = append(javaArgument, "-Dminecraft.api.account.host=https://invalid.invalid")
	javaArgument = append(javaArgument, "-Dminecraft.api.session.host=https://invalid.invalid")
	javaArgument = append(javaArgument, "-Dminecraft.api.services.host=https://invalid.invalid")
	javaArgument = append(javaArgument, "-XX:HeapDumpPath=MojangTricksIntelDriversForPerformance_javaw.exe_minecraft.exe.heapdump")

	for _, value := range JVMArgRaw {
		// Ignore Rules
		if _, ok := value.(map[string]interface{}); ok {
			continue
		}

		finalString := os.ExpandEnv(value.(string))
		javaArgument = append(javaArgument, finalString)
	}

	javaArgument = append(javaArgument, mainClass)

	return javaArgument
}

func GetGameArg(gameArgRaw []interface{}, instance model.Instance, configModel model.Config) []string {
	var gameArgument []string

	usernameQuoted := fmt.Sprintf(`"%s"`, configModel.PlayerName)
	versionNameQuoted := fmt.Sprintf(`"%s"`, instance.Version)
	gameDirQuoted := filepath.Join(config.InstanceDir, instance.Name)
	assetDirQuoted := config.AssetDir

	os.Setenv("AUTH_PLAYER_NAME", usernameQuoted)
	os.Setenv("VERSION_NAME", versionNameQuoted)
	os.Setenv("GAME_DIRECTORY", gameDirQuoted)
	os.Setenv("ASSETS_ROOT", assetDirQuoted)
	os.Setenv("GAME_ASSETS", assetDirQuoted)
	os.Setenv("ASSETS_INDEX_NAME", instance.AssetIndex)
	os.Setenv("LAUNCHER_VERSION", config.LauncherVersion)

	os.Setenv("auth_uuid", `""`)
	os.Setenv("auth_access_token", `""`)
	os.Setenv("clientid", `""`)
	os.Setenv("auth_xuid", `""`)
	os.Setenv("user_type", `""`)
	os.Setenv("version_type", `""`)

	for _, value := range gameArgRaw {
		// Males ngurusin rules
		if _, ok := value.(map[string]interface{}); ok {
			continue
		}

		finalString := os.ExpandEnv(value.(string))
		gameArgument = append(gameArgument, finalString)

	}
	

	return gameArgument
}