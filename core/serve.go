package main

import (
	"satscript/core/cliserver"
	"satscript/core/config"
	"satscript/core/db"
	"satscript/core/log"
	"satscript/core/rpcserver"
	"satscript/core/vm"
	"time"
)

func serverNode(conf *config.CoreConfigs, database *db.NodeDatabase, climan *cliserver.CLIServer, rpcsrv *rpcserver.RPCServer, vmman *vm.NodeScriptVMManager) error {
	// Schließt alle Prozesse in korrekter reihenfolge
	defer rpcsrv.Close()
	defer climan.Close()
	defer vmman.Close()
	defer database.Close()

	log.NODE_PRINTLN("Complete started")

	// Die Schleife wird solange ausgeführt, bis etwas anderes Siganlisiert wird
	for {
		time.Sleep(1 * time.Millisecond)
	}
}
