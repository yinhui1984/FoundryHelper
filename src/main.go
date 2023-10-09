package main

import (
	"flag"
	"foundryhelper/src/configration"
	"foundryhelper/src/features"
	"foundryhelper/src/tui/menu"
	"foundryhelper/src/utils"
	"path"
	"path/filepath"
)

func parseArgs() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "", "config file path")

	var resoureceDirPath string
	flag.StringVar(&resoureceDirPath, "resource", "", "resource directory path")
	flag.Parse()

	if configFilePath == "" {
		configFilePath = path.Join(utils.GetAssemblyFolder(), "config.json")
	}
	configration.LoadConfig(configFilePath)

	if resoureceDirPath == "" {
		resoureceDirPath = path.Join(utils.GetAssemblyFolder(), "../resources")
	}
	if !path.IsAbs(resoureceDirPath) {
		absPath, err := filepath.Abs(resoureceDirPath)
		if err != nil {
			utils.LogFatal("Convert relative config file path to absolute path failed: ", err)
		}
		resoureceDirPath = absPath
	}
	configration.AppConfigInstance.ResourceDir = resoureceDirPath
	utils.LogOK("Resource directory path: ", resoureceDirPath)
}

func main() {
	parseArgs()

	allFeatures := features.GetAllFeatures()
	menu.MakeMenuList(features.GetAllFeaturesString(), "SELECT A FEATURE")

	index := menu.LastChoiseIndex
	if index >= 0 && index < len(allFeatures) {
		f := allFeatures[index]
		f.Invoke()
	}

	utils.LogInfo("Bye!")
}
