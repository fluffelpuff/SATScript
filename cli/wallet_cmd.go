package main

import (
	"fmt"
	"net/rpc"
)

// Wird ausgeführt wenn es sich um einen Wallet befehl handelt
func enterWalletCommand(data []string, conn *rpc.Client) error {
	switch data[0] {
	case "listunspent":
		return listContracts(conn)
	case "listunconfirmet":
		return listContracts(conn)
	case "totalamount":
		return listContracts(conn)
	case "peggs":
		return listContracts(conn)
	case "unspendable":
		return listContracts(conn)
	default:
		return fmt.Errorf("Unbekannter parameter " + data[0])
	}
}
