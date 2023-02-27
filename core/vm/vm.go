package vm

import (
	"satscript/core/config"
	"satscript/core/contract"
	"satscript/core/db"
	"satscript/core/log"
	"strconv"
)

// Stellt den Script VM Manager dar
type NodeScriptVMManager struct {
	database  *db.NodeDatabase
	contracts []*contract.Contract
}

// Signalisiert dem VM-Objekt dass es alle Contracts laden soll und mit der Arbeit beginnen soll
func (obj *NodeScriptVMManager) Start() error {
	// Es werden alle Verfügbaren Contracts aus den Datenbank abgerufen und gestartet
	log.NODE_PRINTLN("All available contracts are retrieved from the database and started")
	_, err := obj.database.FetchAllContractsFromDisk(true)
	if err != nil {
		return err
	}

	// Der Vorgang wurde erfolgreich druchgeführt
	return nil
}

// Signalisiert der VM dass alle Contracts beendet und entladen werden sollen
func (obj *NodeScriptVMManager) Close() error {
	log.NODE_PRINTLN("All contracts are terminated and discharged")
	log.NODE_PRINTLN("All contracts have been terminated and unloaded. TOTAL = " + strconv.Itoa(0))
	return nil
}

// Gibt alle Contracts aus, welche verfügbar sind
func (obj *NodeScriptVMManager) GetAllContracts() ([]*contract.Contract, error) {
	return obj.contracts, nil
}

// Es wird versucht einen neuen Skript Manager zu erstellen
func NewScriptVMManager(conf *config.CoreConfigs, database *db.NodeDatabase) (NodeScriptVMManager, error) {
	return NodeScriptVMManager{database: database}, nil
}
