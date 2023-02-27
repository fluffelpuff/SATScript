package main

import (
	"fmt"
	"net/rpc"
	"os"
	"satscript/core/cliserver"

	"github.com/jedib0t/go-pretty/v6/table"
)

// Printed eine Liste mit allen Contracts
func listContracts(conn *rpc.Client) error {
	// Es werden alle Verfügbaren Contracts abgerufen
	var reply cliserver.ContractListAllResult
	if err := conn.Call("Contracts.ListAll", cliserver.EmptyData{}, &reply); err != nil {
		fmt.Printf("Error:1 user.GetUsers() %+v", err)
	}

	// Es wird geprüft ob mindestens 1 Contract empfangen wurde
	if len(reply.Contracts) < 1 {
		fmt.Println("There are no contracts. Enter 'satscli contract help' for more information")
		fmt.Println("Or visit us at https://github.com/ to report a bug.")
		return nil
	}

	// Die Informationen der Contract werden angezeigt
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Contract ID", "Total Amount", "State", "Total peers"})
	for h, item := range reply.Contracts {
		t.AppendRow([]interface{}{h + 1, item.ContractID, BTC_UINT_FORMATER(item.Amount), item.State, item.TotalPeers})
		if h < len(reply.Contracts) {
			t.AppendSeparator()
		}
	}
	t.Render()

	// Die Contracts wurden erfolgreich abgerufen
	return nil
}

// Wird ausgeführt wenn es sich um einen List befehl handelt
func enterListCommand(data []string, conn *rpc.Client) error {
	switch data[0] {
	case "contracts":
		return listContracts(conn)
	case "wallets":
		return listContracts(conn)
	case "peers":
		return listContracts(conn)
	case "logs":
		return listContracts(conn)
	case "history":
		return listContracts(conn)
	case "containers":
		return listContracts(conn)
	case "bridges":
		return listContracts(conn)
	case "onions":
		return listContracts(conn)
	case "client-processes":
		return listContracts(conn)
	default:
		return fmt.Errorf("Unbekannter parameter " + data[0])
	}
}
