package main

import (
	"fmt"
	"net/rpc"
)

// Wird ausgeführt wenn es sich um einen Contract befehl handelt
func enterContractCommand(data []string, conn *rpc.Client) error {
	switch data[0] {
	case "install":
		return listContracts(conn)
	case "install-from":
		return listContracts(conn)
	case "uninstall":
		return listContracts(conn)
	case "shell":
		return listContracts(conn)
	case "stop":
		return listContracts(conn)
	case "start":
		return listContracts(conn)
	case "disable":
		return listContracts(conn)
	case "enable":
		return listContracts(conn)
	case "update_db":
		return listContracts(conn)
	case "pegin":
		return listContracts(conn)
	case "peginfnc":
		return listContracts(conn)
	case "peers":
		return listContracts(conn)
	case "storage":
		return listContracts(conn)
	case "wallet":
		return listContracts(conn)
	case "ermergancy":
		return listContracts(conn)
	case "debug":
		return listContracts(conn)
	case "help":
		return listContracts(conn)
	case "pegaddresses":
		return listContracts(conn)
	case "inputs":
		return listContracts(conn)
	case "outputs":
		return listContracts(conn)
	case "pfunctions":
		return listContracts(conn)
	case "pvars":
		return listContracts(conn)
	case "accounts":
		return listContracts(conn)
	case "secstate":
		return listContracts(conn)
	case "monitor":
		return listContracts(conn)
	default:
		return fmt.Errorf("Unbekannter parameter " + data[0])
	}
}
