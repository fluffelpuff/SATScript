package rpcserver

import (
	"satscript/core/config"
	"satscript/core/db"
)

type RPCServer struct {
}

// Beendet den RPC Server
func (obj *RPCServer) Close() error {
	return nil
}

// Erstellt einen RPC Server
func CreateNewRPCServer(conf *config.PathConfigs, database *db.NodeDatabase) (RPCServer, error) {
	return RPCServer{}, nil
}
