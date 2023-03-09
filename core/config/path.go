package config

import (
	"fmt"
	"runtime"
)

// Ermittelt die Systemabhänigen Pfade
func DeterminePath() (PathConfigs, error) {
	if runtime.GOOS == "linux" {
		return PathConfigs{
				GeneralSettingsPath: "/etc/satscript.conf",
				NoneRootSocketPath:  "/tmp/ssvmclinr",
				RootSocketPath:      "/tmp/ssvmclroot",
				WalletPath:          "/home/fluffel/wallets",
				DatabasePath:        "/home/fluffelbuff/ssvm/SATScript",
			},
			nil
	} else if runtime.GOOS == "darwin" {
		return PathConfigs{
				GeneralSettingsPath: "/Library/SATScript/satscript.conf",
				NoneRootSocketPath:  "/Library/SATScript/nrootapi",
				RootSocketPath:      "/Library/SATScript/rootapi",
				WalletPath:          "/Library/SATScript/wdb",
				DatabasePath:        "/Library/SATScript/db",
			},
			nil
	} else {
		return PathConfigs{}, fmt.Errorf("determine_paths: Unsupported host os")
	}
}
