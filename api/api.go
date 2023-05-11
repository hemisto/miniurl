package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/exp/slog"
)

type Handler interface {
	Hash(url string) (hash string, err error)
}

type API struct {
	handler Handler
}

func Bind(r *httprouter.Router, h Handler) {
	a := &API{handler: h}
	r.POST("/api/v1/url", a.AddUrl)
}

type AddUrlReq struct {
	Url string `json:"url"`
}

type AddUrlResp struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

type ErrResp struct {
	Msg string `json:"msg"`
}

func (a *API) AddUrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var v AddUrlReq
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		//ToDo
		return
	}

	hash, err := a.handler.Hash(v.Url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(ErrResp{Msg: "error occured while generating hash"})
		if err != nil {
			slog.Error(err.Error())
		}
		return
	}

	err = json.NewEncoder(w).Encode(AddUrlResp{Url: v.Url, Hash: hash})
	if err != nil {
		slog.Error(err.Error())
	}
}
