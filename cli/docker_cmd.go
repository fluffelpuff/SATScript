package main

import (
	"fmt"
	"net/rpc"
)

// Wird ausgeführt wenn es sich um einen Docker befehl handelt
func dockerCommand(date []string, conn *rpc.Client) error {
	fmt.Println("The function you are running is not ready yet. Try again after an update")
	return nil
}
