package handlers

import (
	"encoding/json"
	"github.com/auctionee/goods/pkg/data"
	"github.com/auctionee/goods/pkg/db"
	"github.com/auctionee/goods/pkg/logger"
	"io/ioutil"
	"net/http"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.GetLogger(r.Context())
	creds := data.Credentials{}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.Println("cant parse request body in /info")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, creds)
	if err != nil {
		l.Println("cant unmarshal request body in /info")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	goods := db.GetGoods(creds)
	resp, err := json.Marshal(goods)
	if err != nil {
		l.Println("cant marshal result in /info")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		l.Info("error writing response in /info: ", err)
	}
}
