package features

import (
	install "foundryhelper/src/features/install"
	new_project "foundryhelper/src/features/new_project"
)

type Feature interface {
	GetFeatureName() string
	Invoke()
}

var featuresList = make([]Feature, 0)

func init() {
	InitFeatures()
}

func InitFeatures() {
	featuresList = append(featuresList, new_project.NewProject{FeatureName: "创建新项目"})
	featuresList = append(featuresList, install.Install{FeatureName: "安装第三方包"})
}

func GetAllFeatures() []Feature {
	return featuresList
}

func GetAllFeaturesString() []string {
	featuresStringList := make([]string, 0)
	for _, feature := range featuresList {
		featuresStringList = append(featuresStringList, feature.GetFeatureName())
	}
	return featuresStringList
}
