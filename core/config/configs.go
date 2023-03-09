package config

import (
	"runtime"
)

type PathConfigs struct {
	GeneralSettingsPath string
	NoneRootSocketPath  string
	RootSocketPath      string
	WalletPath          string
	DatabasePath        string
}

func (obj *PathConfigs) GetDatabaseViewFilePath() (string, error) {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		return obj.DatabasePath + "/master.db", nil
	}
	return "", nil
}

type ConfigFile struct {
}
