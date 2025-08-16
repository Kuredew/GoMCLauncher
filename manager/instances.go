package manager

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Kuredew/GoMCLauncher/config"
	managerpanel "github.com/Kuredew/GoMCLauncher/manager/manager_panel"
	managerutils "github.com/Kuredew/GoMCLauncher/manager/manager_utils"
	"github.com/Kuredew/GoMCLauncher/model"
	"github.com/Kuredew/GoMCLauncher/services"
	"github.com/Kuredew/GoMCLauncher/utils"
)

var Argument []string
var gameDir string

func getDependency(instance model.Instance) (map[string]interface{}, map[string]interface{}, string) {
	
	var dependencyInfo map[string]interface{}
	var assetList map[string]interface{}
	var assetIndex string


	versionManifest := services.GetVersionManifest()
	versionList := versionManifest["versions"].([]interface{})
	
	// search version in version manifest
	for _, value := range versionList {
		id := value.(map[string]interface{})["id"].(string)

		if id == instance.Version {
			dependencyInfo, assetList, assetIndex = services.GetDependency(value.(map[string]interface{}))
		}
	}

	return dependencyInfo, assetList, assetIndex
}

func getJavaRuntime() string {
	log.Print("Getting java-runtime...")
	var javaExec string
	var javaDownloadUrl string

	osStr := utils.GetOSStr()
	switch osStr {
		case "windows":
			javaDownloadUrl = config.JavaWinDownloadUrl
			javaExec = "java.exe"
		case "linux":
			javaDownloadUrl = config.JavaLinuxDownloadUrl
			javaExec = "java"
		default:
			javaDownloadUrl = config.JavaMacDownloadUrl
			javaExec = "java"
	}

	for {
		// search jdk folder from java-runtime directory
		items, _ := os.ReadDir(config.JavaRuntimeDir)
		for _, item := range items {
			itemName := item.Name()

			if !strings.Contains(itemName, "jdk") {
				continue
			}

			JavaRuntimeFile := filepath.Join(config.JavaRuntimeDir, itemName, "bin", javaExec)

			if !utils.FileExists(JavaRuntimeFile) {
				log.Printf("%s Not Exist", JavaRuntimeFile)

				break
			}

			log.Printf("Launching Java at %s", JavaRuntimeFile)
			return JavaRuntimeFile
		}
		
		// if zip not exist, download it!
		if !utils.FileExists(config.JavaRuntimeZip) {
			log.Print("Java Archive not exist")
			utils.Download(config.JavaRuntimeZip, javaDownloadUrl)
		}

		log.Print("Extracting Java...")
		err := utils.ExtractZIP(config.JavaRuntimeDir, config.JavaRuntimeZip)
		if err != nil {
			log.Printf("Extracting java error, please wait...")
			utils.Download(config.JavaRuntimeZip, javaDownloadUrl)
		}
	}
}

func StartInstance(instance model.Instance, configModel model.Config) error {
	cooldown := 3

	// Mwehehe
	for {
		fmt.Printf("Launching in %v second...\r", cooldown)
		cooldown--
		time.Sleep(1 * time.Second)
		
		if cooldown == 0 {
			break
		}
	}

	gameDir = filepath.Join(config.InstanceDir, instance.Name)
	os.MkdirAll(gameDir, 0755)

	// Get Dependency
	dependencyInfo, assetList, assetIndex := getDependency(instance)
	instance.AssetIndex = assetIndex
			
	classpath := managerutils.GetLibraries(dependencyInfo)
	Argument = managerutils.GetArg(dependencyInfo, classpath, instance, configModel)
	managerutils.GetAsset(assetList)

	javaPath := getJavaRuntime()
	log.Printf("Launching %s...\n\n", instance.Name)

	utils.ExecuteCMD(javaPath, Argument...)

	log.Print("\nClearing Cache...\n")
	os.RemoveAll(config.NativeLibrariesDir)

	fmt.Print("\n\n[ Type any to return Home ]")
	_ = utils.AskUserInput()

	return nil
}

var ErrBack error = errors.New("back")

func Initialize(configModel model.Config) error {
	for {
		instance, err := managerpanel.InstancePanel()
		
		if err != nil {
			if errors.Is(err, managerpanel.ErrInstancePanelNewInstance) {
				managerpanel.CreateNewInstancePanel()

				continue
			} else if errors.Is(err, managerpanel.ErrInstancePanelBack) {
				return ErrBack
			}
		}
		
LoopMenu:
		for {
			options := []string{"Play", "Modify Instance", "", "Back", "", "Delete Instance [Danger Zone]"}

			headerString := fmt.Sprintf("Select Options for %s", instance.Name)

			userSelected, _ := utils.CreatePanel(headerString, options)

			switch userSelected {
			case 0:
				err := StartInstance(instance, configModel)
				if err != nil {
					return err
				}

				return ErrBack
			case 1:
				managerpanel.ModifyInstancePanel(&instance)
				continue
			case 3:
				break LoopMenu
			case 5:
				err := managerutils.DeleteInstances(instance)
				if err != nil {
					continue
				} else {
					break LoopMenu
				}
			}
		}
	}
}