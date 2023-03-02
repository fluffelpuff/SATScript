package main

import (
	"fmt"
	"net/rpc"
	"os"
)

func InitConn() (*rpc.Client, error) {
	// Es wird versucht eine Verbindung zum UnixSocket herzustellrn
	conn, err := rpc.DialHTTP("unix", "path")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn, nil
}

func main() {
	// Die Parameter werden eingelesen
	if len(os.Args) == 1 {
		fmt.Println("Bitte gib einen befehl an")
		return
	}
	readed_var := os.Args[1:]

	// Es wird geprüft um welchen befehl es sich handelt
	if readed_var[0] == "list" {
		conn, err := InitConn()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := enterListCommand(readed_var[1:], conn); err != nil {
			fmt.Println(err)
		}
	} else if readed_var[0] == "contract" {
		conn, err := InitConn()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := enterContractCommand(readed_var[1:], conn); err != nil {
			fmt.Println(err)
		}
	} else if readed_var[0] == "getinfo" || readed_var[0] == "info" {
		conn, err := InitConn()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := getInfoCommand(readed_var[1:], conn); err != nil {
			fmt.Println(err)
		}
	} else if readed_var[0] == "dashboard" {
		conn, err := InitConn()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := startDashboard(readed_var[1:], conn); err != nil {
			fmt.Println(err)
		}
	} else if readed_var[0] == "script" {
		if err := scriptCommand(readed_var[1:]); err != nil {
			fmt.Println(err)
		}
	} else if readed_var[0] == "docker" {
		conn, err := InitConn()
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := dockerCommand(readed_var[1:], conn); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Unbekannter befehl: " + readed_var[0])
		return
	}
}
