package service

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"fmt"
)

func Route(router *mux.Router, service Service) *mux.Route {
	fmt.Println("Creating route: " + service.Path())
	return router.Handle(service.Path(), createRpcService(service))
}

func createRpcService(service Service) *rpc.Server {
	rpcServer := rpc.NewServer()

	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	rpcServer.RegisterService(service, service.Name())

	return rpcServer
}