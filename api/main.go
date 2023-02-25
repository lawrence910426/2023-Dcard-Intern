package main

import (
	"net/http"

	routes "dcard-intern/router"
)

func main() {

	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)

}
