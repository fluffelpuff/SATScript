package main

import (
	"fmt"
	"net/rpc"
	"satvm/core/cliserver"
	"strconv"
)

// Wird ausgeführt wenn es sich um einen GetInfo befehl handelt
func getInfoCommand(date []string, conn *rpc.Client) error {
	// Es wird geprüft das kein weitere befehl auf dem Commandstack liegt
	if len(date) != 0 {
		return fmt.Errorf("To many arguments for getinfo")
	}

	// Es werden alle Verfügbaren Contracts abgerufen
	var reply cliserver.VmProcessInfo
	if err := conn.Call("SatScriptProcess.GetInfo", cliserver.EmptyData{}, &reply); err != nil {
		fmt.Printf("Error:1 user.GetUsers() %+v", err)
	}

	// Die Informationen werden angezeigt
	fmt.Println("Master public key                 = ")
	fmt.Println("Master onion address              = ")
	fmt.Println("Network usage                     = 0.00% / 0.00 Mbit/s")
	fmt.Println("CPU Usage                         = 0.01%")
	fmt.Println("RAM Usage                         = 10 MB")
	fmt.Println("Run as Root                       = Yes")
	fmt.Println("Total containers                  = 0")
	fmt.Println("Total contracts                   = " + strconv.Itoa(int(reply.TotalContracts)))
	fmt.Println("Total peers                       = " + strconv.Itoa(int(reply.TotalPeers)))
	fmt.Println("Bitcoin                           = Core (v.023)")
	fmt.Println("  > Peers                         = 0")
	fmt.Println("  > Hight                         = 0")
	fmt.Println("Coin network                      = TestNet")
	fmt.Println("Total storage used                = 0.00 MB")
	fmt.Println("SatScript version                 = go v" + strconv.Itoa(int(reply.Version)))
	fmt.Println("SatScript container version       = go v0.0.1")
	fmt.Println("SatScript cli-tool version        = go v0.0.1")
	fmt.Println("SatScript wallet firmware         = 0.0.1")
	fmt.Println("Web interface (HTTP / WS)         = Disabled (default)")
	fmt.Println("Docker                            = Not supported")
	fmt.Println("Tor mode                          = Integrated (bine-go v0.0.1)")
	fmt.Println("Total Pegget unlocked amount      = 0.00000000 BTC")
	fmt.Println("Total Pegget locked amount        = 0.00000000 BTC")

	// Die Contracts wurden erfolgreich abgerufen
	return nil
}
