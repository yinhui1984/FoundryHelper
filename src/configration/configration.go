package configration

import (
	"foundryhelper/src/utils"
	"path"
	"path/filepath"
)

type RpcEndPoint struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Configration struct {
	RpcEndPoints []RpcEndPoint `json:"rpc_end_points"`
	ResourceDir  string        `json:"-"`
	DefaultIDE   string        `json:"default_ide"`
}

var AppConfigInstance Configration = Configration{}

func LoadConfig(configFilePath string) {
	// if configFilePath is relative path, convert it to absolute path
	if !path.IsAbs(configFilePath) {
		absPath, err := filepath.Abs(configFilePath)
		if err != nil {
			utils.LogFatal("Convert relative config file path to absolute path failed: ", err)
		}
		configFilePath = absPath
	}

	err := utils.LoadJSONConfig(configFilePath, &AppConfigInstance)
	if err != nil {
		utils.LogFatal("Load config file failed: ", err)
	} else {
		utils.LogOK("Load app config file success: ", configFilePath)
	}
}
