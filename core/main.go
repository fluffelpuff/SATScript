package main

import (
	"fmt"
	"runtime"
	"satscript/core/cliserver"
	"satscript/core/config"
	"satscript/core/db"
	"satscript/core/log"
	"satscript/core/rpcserver"
	"satscript/core/vm"
)

func determine_paths() (config.CoreConfigs, error) {
	// Es wird geprüft ob es sich um ein Linux oder BSD basiertes System handelt
	if runtime.GOOS == "linux" {
		//return config.CoreConfigs{GeneralSettingsPath: "/tmp/ssvm", NoneRootSocketPath: "/etc/satscript.conf", WalletPath: "/home/fluffel/wallets", DatabasePath: "/var/lib/SATScript"}, nil
		return config.CoreConfigs{GeneralSettingsPath: "/etc/satscript.conf", NoneRootSocketPath: "/tmp/ssvmclinr", RootSocketPath: "/tmp/ssvmclroot", WalletPath: "/home/fluffel/wallets", DatabasePath: "/home/fluffelbuff/ssvm/SATScript"}, nil
	} else {
		return config.CoreConfigs{}, fmt.Errorf("Unsupported host os")
	}
}

func main() {
	// Es wird versucht die Passenden Pfade für SATScriptd zu ermitteln
	paths, err := determine_paths()
	if err != nil {
		log.NODE_EPRINTLN(err)
		return
	}

	// Die Datenbank wird erzeugt und oder geladen
	db, err := db.LoadDatabase(&paths, true)
	if err != nil {
		log.NODE_EPRINTLN(err)
		return
	}

	// Die VM wird erzeugt
	vm_manager, err := vm.NewScriptVMManager(&paths, &db)
	if err != nil {
		log.NODE_EPRINTLN(err)
		return
	}

	// Der VM-Manager wird gestartet
	if err := vm_manager.Start(); err != nil {
		log.NODE_EPRINTLN(err)
		return
	}

	// Der CLI Server wird erzeugt
	cli_control, err := cliserver.NewCLIServer(&paths, &db, &vm_manager)
	if err != nil {
		log.NODE_EPRINTLN(err)
		return
	}

	// Der RPC Server wird erzeugt
	rpc_server, err := rpcserver.CreateNewRPCServer(&paths, &db)
	if err != nil {
		log.NODE_EPRINTLN(err)
		return
	}

	// Diese Funktion wird ausgeführt und hält den Prozess geöffnet
	if err := serverNode(&paths, &db, &cli_control, &rpc_server, &vm_manager); err != nil {
		log.NODE_EPRINTLN(err)
		return
	}
}
