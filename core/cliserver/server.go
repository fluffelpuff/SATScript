package cliserver

import (
	"net"
	"net/http"
	"net/rpc"
	"os"
	"satscript/core/config"
	"satscript/core/contract"
	"satscript/core/db"
	"satscript/core/info"
	"satscript/core/log"
	"satscript/core/vm"
	"strconv"
)

type EmptyData struct{}

type Contracts struct {
	SS_DB  *db.NodeDatabase
	VM_MAN *vm.NodeScriptVMManager
}

type ContractListAllResult struct {
	Contracts []contract.ContractInformations
}

type VmProcessInfo struct {
	UsedStorage        string
	PublicOnionAddress string
	Version            uint
	TotalContracts     uint
	TotalPeers         int
}

type SatScriptProcess struct {
	SS_DB  *db.NodeDatabase
	VM_MAN *vm.NodeScriptVMManager
}

type ScriptCompileInput struct {
}

type ScriptCompileResult struct {
}

type ScriptIo struct {
	SS_DB  *db.NodeDatabase
	VM_MAN *vm.NodeScriptVMManager
}

// Gibt die Wichtisten Informationen zurück
func (obj *SatScriptProcess) GetInfo(_ EmptyData, result *VmProcessInfo) error {
	// Es werden alle Contracts abgerufen
	all_contracts, err := obj.VM_MAN.GetAllContracts()
	if err != nil {
		return err
	}

	// Die Gesamtzahl aller Peers wird ermittelt
	total_conn := uint(0)
	for _, item := range all_contracts {
		total, err := item.GetTotalPeerCount()
		if err != nil {
			return err
		}
		total_conn += total
	}

	// Die Antwort wird gebaut
	result.UsedStorage = "0x00"
	result.TotalContracts = uint(len(all_contracts))
	result.TotalPeers = int(total_conn)
	result.Version = info.VERSION

	// Der Vorgang wurde erfolgreich druchgeführt
	log.NODE_PRINTLN("Transfer GetInfo data")
	return nil
}

// Gibt eine Liste mit allen geladenen Smart Contracts zurück
func (rh *Contracts) ListAll(_ EmptyData, reply *ContractListAllResult) error {
	// Es werden alle Contracts abgerufen
	contr, err := rh.VM_MAN.GetAllContracts()
	if err != nil {
		return err
	}

	// Die Contract Informationen werden übermittelt
	for _, item := range contr {
		info_obj, err := item.GetContractInformation()
		if err != nil {
			return err
		}
		reply.Contracts = append(reply.Contracts, *info_obj)
	}

	// Der Vorgang wurde erfolgreich ausgeführt
	log.NODE_PRINTLN(strconv.Itoa(len(contr)) + " contract data was transmitted")
	return nil
}

// Startet einen neuen CLI-Server
func newCLIServer(path string, root_privs bool, db *db.NodeDatabase, vmobj *vm.NodeScriptVMManager) (*Contracts, error) {
	// Es wird geprüft ob die Socket Datei vorhanden ist
	if _, err := os.Stat(path); err == nil {
		e := os.Remove(path)
		if e != nil {
			return &Contracts{}, err
		}
	}

	// Es wird versucht einen neuen UNIX Socket zu erstellen
	socket, err := net.Listen("unix", path)
	if err != nil {
		return &Contracts{}, err
	}

	// Der RPC Server wird erstellt
	contract_rpc := new(Contracts)
	process_rpc := new(SatScriptProcess)
	script_io := new(SatScriptProcess)
	contract_rpc.SS_DB = db
	contract_rpc.VM_MAN = vmobj
	process_rpc.SS_DB = db
	process_rpc.VM_MAN = vmobj
	script_io.SS_DB = db
	script_io.VM_MAN = vmobj
	rpc.Register(contract_rpc)
	rpc.Register(process_rpc)
	rpc.Register(script_io)
	rpc.HandleHTTP()

	// Der Webserver wird erstellt
	go http.Serve(socket, nil)

	// Es wird versucht einen
	if root_privs {
		log.NODE_PRINTLN("The CLI server has been started. ROOT = yes, PATH = " + path)
	} else {
		log.NODE_PRINTLN("The CLI server has been started. ROOT = no, PATH = " + path)
	}
	return contract_rpc, nil
}

// Stellt ein CLI-Server Manager da
type CLIServer struct {
	server []*Contracts
}

// Beendet den CLI Server
func (obj *CLIServer) Close() error {
	return nil
}

// Erzeugt einen neuen CLI Server
func NewCLIServer(config *config.PathConfigs, db *db.NodeDatabase, vmobj *vm.NodeScriptVMManager) (CLIServer, error) {
	// Der Normale RPC Server wird erstellt, keine Administratorrechte
	non_root_socket, err := newCLIServer(config.NoneRootSocketPath, false, db, vmobj)
	if err != nil {
		return CLIServer{}, err
	}

	// Es wird geprüft ob der Aktuelle Benutzer als Root unterwegs ist
	if currentIsPrivUser(config) {
		// Der Root RPC Server wird erstellt, Administratorrechte
		root_socket, err := newCLIServer(config.RootSocketPath, true, db, vmobj)
		if err != nil {
			log.NODE_PRINTLN("The root RPC socket could not be created, you do not have the required rights.")
		}

		// Der Vorgang wurde erfolgreich fertigestellt
		return CLIServer{server: []*Contracts{non_root_socket, root_socket}}, nil
	}

	// Der Vorgang wurde erfolgreich fertigestellt
	return CLIServer{server: []*Contracts{non_root_socket}}, nil
}
