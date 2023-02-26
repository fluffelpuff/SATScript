package config

import "runtime"

type CoreConfigs struct {
	GeneralSettingsPath string
	NoneRootSocketPath  string
	RootSocketPath      string
	WalletPath          string
	DatabasePath        string
}

func (obj *CoreConfigs) GetDatabaseViewFilePath() (string, error) {
	if runtime.GOOS == "linux" {
		return obj.DatabasePath + "/master.db", nil
	}
	return "", nil
}
