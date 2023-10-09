package newproject

import (
	"foundryhelper/src/configration"
	"foundryhelper/src/tui/input"
	"foundryhelper/src/tui/menu"
	"foundryhelper/src/utils"
	"path"
)

type NewProject struct {
	FeatureName string
}

func (n NewProject) GetFeatureName() string {
	return n.FeatureName
}

func (n NewProject) Invoke() {

	rpcList := configration.AppConfigInstance.RpcEndPoints
	repcListForMenu := make([]string, 0)
	for _, rpc := range rpcList {
		repcListForMenu = append(repcListForMenu, "[ "+utils.MakeTextCenter(rpc.Name, 11)+" ] "+rpc.Url)
	}

	menu.MakeMenuList(repcListForMenu, "SELECT A RPC ENDPOINT")

	index := menu.LastChoiseIndex
	if index < 0 || index >= len(rpcList) {
		utils.LogError("Invalid RPC Endpoint index: ", index)
		return
	}

	rpcSelected := rpcList[index]
	utils.LogOK("Selected RPC Endpoint: ", rpcSelected.Name, " ", rpcSelected.Url)

	createNewProject(rpcSelected)
}

func createNewProject(rpcEndpoint configration.RpcEndPoint) {

	input.MakeInputs([]string{"Project Name"})
	projectName := input.LastFielsValues[0]
	if projectName == "" {
		projectName = "NewProject"
	}

	resoureDir := configration.AppConfigInstance.ResourceDir
	templateDir := path.Join(resoureDir, "ProjectTemplate")
	targetDir := path.Join(utils.GetWorkingFolder(), projectName)

	utils.LogInfo("copy template from ", templateDir, " to ", targetDir)

	if utils.IsFolderOrFileExist(targetDir) &&
		!utils.AskForYes("Target directory already exist, overwrite it?") {
		return
	}

	_, err := utils.RunCommandLine("mkdir -p " + targetDir)
	if err != nil {
		utils.LogError("Create target directory failed: ", err)
		return
	}
	out, err := utils.RunCommandLine("cp -rvf " + templateDir + "/ " + targetDir)
	if err != nil {
		utils.LogError("Copy template failed: ", err, out)
		return
	}

	err = utils.ReplaceFirstInFile("theNet\\s*=.*\n", "theNet = \""+rpcEndpoint.Url+"\"\n", targetDir+"/foundry.toml")
	if err != nil {
		utils.LogError("Replace RPC Endpoint failed: ", err)
		return
	}

	out, err = utils.RunCommandLine2(configration.AppConfigInstance.DefaultShell, "cd "+targetDir+" && forge test -vvv")
	if err != nil {
		utils.LogError("Forge test failed: ", err, out)
		return
	} else {
		utils.LogOK("Forge test success", out)
	}

	out, err = utils.RunCommandLine(configration.AppConfigInstance.DefaultIDE + " " + targetDir)
	if err != nil {
		utils.LogError("open IDE failed: ", err, out)
		return
	} else {
		utils.LogOK("open IDE success", out)
	}
	utils.LogInfo("New project created successfully: ", targetDir)
}

func (n NewProject) String() string {
	return n.FeatureName
}
