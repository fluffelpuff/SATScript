package main

import (
	"fmt"
	"net/rpc"
	"os"
)

func main() {
	// Die Parameter werden eingelesen
	if len(os.Args) == 1 {
		fmt.Println("Bitte gib einen befehl an")
		return
	}
	readed_var := os.Args[1:]

	// Es wird versucht eine Verbindung zum UnixSocket herzustellrn
	conn, err := rpc.DialHTTP("unix", "/tmp/ssvmclinr")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Es wird geprüft um welchen befehl es sich handelt
	if readed_var[0] == "list" {
		if err := enterListCommand(readed_var[1:], conn); err != nil {
			fmt.Println(err)
		}
	} else if readed_var[0] == "contract" {
		if err := enterContractCommand(readed_var[1:], conn); err != nil {
			fmt.Println(err)
		}
	} else if readed_var[0] == "getinfo" || readed_var[0] == "info" {
		if err := getInfoCommand(readed_var[1:], conn); err != nil {
			fmt.Println(err)
		}
	} else if readed_var[0] == "dashboard" {
		if err := startDashboard(readed_var[1:], conn); err != nil {
			fmt.Println(err)
		}
	} else if readed_var[0] == "script" {
		if err := scriptCommand(readed_var[1:], conn); err != nil {
			fmt.Println(err)
		}
	} else if readed_var[0] == "docker" {
		if err := dockerCommand(readed_var[1:], conn); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Unbekannter befehl: " + readed_var[0])
		if e := conn.Close(); e != nil {
			fmt.Println(e)
		}
		return
	}

	// Die Verbindung wird geschlossen
	if e := conn.Close(); e != nil {
		fmt.Println(e)
	}
}
