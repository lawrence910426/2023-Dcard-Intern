package main

import (
	rpc_server "dcard-intern/proto"
	routes "dcard-intern/router"
	"net/http"
)

func main() {
	go rpc_server.StartServer()

	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)
}
