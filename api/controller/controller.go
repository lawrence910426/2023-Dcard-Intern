package controller

import (
	"github.com/gorilla/mux"

	"net/http"

	"dcard-intern/services"
)

type HeadNode struct {
	NextPageKey string
}

type Node struct {
	Article     string
	NextPageKey string
}

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

func GetHead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"]

	var listHead HeadNode
	listHead.NextPageKey = queryId

	response := ApiResponse{"200", listHead}
	services.ResponseWithJson(w, http.StatusOK, response)
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"]

	var listNode Node
	listNode.NextPageKey = queryId

	response := ApiResponse{"200", listNode}
	services.ResponseWithJson(w, http.StatusOK, response)
}
