package main

import (
	"context"
	"net/http"
	"time"

	"github.com/cretz/bine/tor"
)

type TorSever struct {
	OnionURL string
	OnionObj *tor.OnionService
}

type NetworkHandler struct {
	TorObj  *tor.Tor
	Servers []TorSever
}

func New() (NetworkHandler, error) {
	t, err := tor.Start(nil, &tor.StartConf{DataDir: "./data"})
	if err != nil {
		return NetworkHandler{}, nil
	}
	defer t.Close()

	return NetworkHandler{TorObj: t}, nil
}

func (obj *NetworkHandler) CreateNewEndPoint() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Dark World!"))
	})

	listenCtx, listenCancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer listenCancel()

	onion, err := obj.TorObj.Listen(listenCtx, &tor.ListenConf{LocalPort: 8080, RemotePorts: []int{80}, Version3: true})
	if err != nil {
		return err
	}

	defer onion.Close()

	new_srv_obj := TorSever{OnionObj: onion, OnionURL: onion.ID}
	obj.Servers = append(obj.Servers, new_srv_obj)

	return http.Serve(onion, nil)
}
