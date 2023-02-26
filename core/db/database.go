package db

import (
	"database/sql"
	"os"
	"satvm/core/config"
	"satvm/core/contract"
	"satvm/core/log"
	"satvm/core/peer"
	"strconv"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

// Dieser SQL Code erstellt alle Tabellen die Notwendig sind
const DB_BASE_TABLES = `
CREATE TABLE "contract_db_map" (
	"map_id"	INTEGER UNIQUE,
	"contract_id"	TEXT,
	"db_id"	TEXT,
	PRIMARY KEY("map_id" AUTOINCREMENT)
);
CREATE TABLE "peers_view" (
	"peer_db_id"	INTEGER UNIQUE,
	"adr_type"	INTEGER,
	"endpoint"	INTEGER,
	"blacklisted"	INTEGER DEFAULT 0,
	PRIMARY KEY("peer_db_id" AUTOINCREMENT)
);
`

// Stellt ein Datenbank Objekt dar
type NodeDatabase struct {
	DbFilePath string
	SQLConn    *sql.DB
	lock       sync.Mutex
}

// Wird verwendet um alle bekannten Peers für ein Contract abzurufen
func (obj *NodeDatabase) GetAllContractPeers(contract_obj *contract.Contract) ([]*peer.NodePeer, error) {
	return []*peer.NodePeer{}, nil
}

// Wird verwendet um einen neuen Contract zu Initalisieren
func (obj *NodeDatabase) InitNewContract(contract_obj *contract.Contract) error {
	return nil
}

// Listet alle bekannten Smart Contracts auf
func (obj *NodeDatabase) FetchAllContractsFromDisk(fetch_full bool) ([]*contract.Contract, error) {
	return_value := []*contract.Contract{}
	log.NODE_PRINTLN("All Contracts have been retrieved. TOTAL = " + strconv.Itoa(len(return_value)))
	return return_value, nil
}

// Wird audgerufen um die Datenbank zu schliefen
func (obj *NodeDatabase) Close() error {
	obj.lock.Lock()
	if obj.SQLConn != nil {
		if err := obj.SQLConn.Close(); err != nil {
			obj.lock.Unlock()
			return err
		}
	}
	obj.lock.Unlock()
	log.NODE_PRINTLN("The database has been closed. PATH = " + obj.DbFilePath)
	return nil
}

// Lädt eine Datenbank
func LoadDatabase(config *config.CoreConfigs, create_new bool) (NodeDatabase, error) {
	// Der Dateipfad für die Masterdatei wird abgerufen
	master_file_path, err := config.GetDatabaseViewFilePath()
	if err != nil {
		return NodeDatabase{}, err
	}

	// Es wird geprüft ob der Entsprechende Ordner vorhanden ist
	var sql_db *sql.DB
	is_a_new_db := false
	if _, err := os.Stat(config.DatabasePath); os.IsNotExist(err) {
		// Es wird versucht den Ordner zu erstellen
		if err := os.Mkdir(config.DatabasePath, os.ModePerm); err != nil {
			return NodeDatabase{}, err
		}

		// Die Masterdatei zum verwalten der Datenbank wird angelegt
		sql_db, err = sql.Open("sqlite3", master_file_path)
		if err != nil {
			return NodeDatabase{}, err
		}
		is_a_new_db = true
	}

	// Es wird geprüft ob es eine Masterdatei in dem Ordner gibt
	if _, err := os.Stat(master_file_path); os.IsNotExist(err) {
		sql_db, err = sql.Open("sqlite3", master_file_path)
		if err != nil {
			return NodeDatabase{}, err
		}
		is_a_new_db = true
	} else {
		// Es wird versucht die SQL Datei zu laden
		sql_db, err = sql.Open("sqlite3", master_file_path)
		if err != nil {
			return NodeDatabase{}, err
		}
	}

	// Es wird geprüft ob es sich um eine neue Datenbank handelt, wenn ja werden alle benötigten Tabellen angelegt
	if is_a_new_db {
		_, err = sql_db.Exec(DB_BASE_TABLES)
		if err != nil {
			return NodeDatabase{}, err
		}
	} else {
		log.NODE_PRINTLN("The database was loaded successfully. DB_PATH = " + config.DatabasePath)
	}

	// Es wird ein neues Datenbank Objekt zurückgegeben
	return NodeDatabase{SQLConn: sql_db, DbFilePath: config.DatabasePath}, nil
}
