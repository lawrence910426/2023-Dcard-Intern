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

type TailNode struct {
	Article string
}

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

func GetHead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"]

	var listHead HeadNode
	listHead.NextPageKey = services.Get(queryId + "_Key")

	response := ApiResponse{"200", listHead}
	services.ResponseWithJson(w, http.StatusOK, response)
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"]

	nextPage := services.Get(queryId + "_Key")
	article := services.Get(queryId + "_Article")

	var response ApiResponse
	if nextPage == "" {
		response = ApiResponse{"200", &TailNode{Article: article}}
	} else {
		response = ApiResponse{"200", &Node{NextPageKey: nextPage, Article: article}}
	}

	services.ResponseWithJson(w, http.StatusOK, response)
}
