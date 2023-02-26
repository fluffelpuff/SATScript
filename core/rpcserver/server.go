package rpcserver

import (
	"satvm/core/config"
	"satvm/core/db"
)

type RPCServer struct {
}

// Beendet den RPC Server
func (obj *RPCServer) Close() error {
	return nil
}

// Erstellt einen RPC Server
func CreateNewRPCServer(conf *config.CoreConfigs, database *db.NodeDatabase) (RPCServer, error) {
	return RPCServer{}, nil
}
