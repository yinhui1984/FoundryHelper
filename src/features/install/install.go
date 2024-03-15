package install

import (
	"foundryhelper/src/configration"
	"foundryhelper/src/tui/menu"
	"foundryhelper/src/utils"
	"os"
	"path"
)

type Install struct {
	FeatureName string
}

func (n Install) GetFeatureName() string {
	return n.FeatureName
}

func (n Install) Invoke() {
	workingDir, err := os.Getwd()
	if err != nil {
		utils.LogError("Get working directory failed: ", err)
		return
	}

	if !utils.IsFoundryProjectFolder(workingDir) {
		utils.LogError("Not a foundry project folder: ", workingDir)
		return
	}

	if !utils.IsGitProjectFolder(workingDir) {
		asnwer := utils.AskForYes(workingDir + " is not a git project folder, initialize it?")
		if asnwer {
			output, err := utils.RunCommandLine2(configration.AppConfigInstance.DefaultShell, "git init "+workingDir)
			if err != nil {
				utils.LogError("Initialize git project failed: ", err)
				return
			} else {
				utils.LogOK("Initialize git project success: ", output)
			}
		}
	}

	installCmds := configration.AppConfigInstance.InstallCmds
	titles := make([]string, 0)
	for _, cmd := range installCmds {
		titles = append(titles, cmd.Title)
	}

	menu.MakeMenuList(titles, "SELECT A PACKAGE TO INSTALL")
	index := menu.LastChoiseIndex
	if index < 0 || index >= len(installCmds) {
		utils.LogError("Invalid install command index: ", index)
		return
	}

	cmd := installCmds[index]
	utils.LogOK("Selected install command: ", cmd.Title, " ", cmd.Command)

	stopWatchChann := make(chan struct{})

	utils.WatchFolder(path.Join(workingDir, "lib/"), stopWatchChann)

	err = utils.RunCommandLine3(configration.AppConfigInstance.DefaultShell, cmd.Command)
	close(stopWatchChann)

	if err != nil {
		utils.LogError("Run install command failed: ", err)
		return
	} else {
		utils.LogOK("Run install command success: ", cmd.Command)
	}

	//remapping
	err = utils.RunCommandLine3(configration.AppConfigInstance.DefaultShell, "forge remappings > remappings.txt")
	if err != nil {
		utils.LogError("Run install command failed: ", err)
		return
	} else {
		utils.LogOK("Run install command success: ", cmd.Command)
	}
}
