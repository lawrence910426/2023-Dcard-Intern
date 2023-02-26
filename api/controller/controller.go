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

	services.ResponseWithJson(w, http.StatusOK, listHead)
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"]

	nextPage := services.Get(queryId + "_Key")
	article := services.Get(queryId + "_Article")

	if nextPage == "" {
		services.ResponseWithJson(w, http.StatusOK, &TailNode{Article: article})
	} else {
		services.ResponseWithJson(w, http.StatusOK,
			&Node{NextPageKey: nextPage, Article: article})
	}
}
