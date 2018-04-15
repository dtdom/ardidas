package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dtdom/ardidas/models"
	"github.com/dtdom/ardidas/utils"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	item.ID = mux.Vars(r)["itemid"]
	item.Get()
	utils.RespondWithJSON(w, http.StatusOK, item)
}

func SellItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	item.ID = mux.Vars(r)["itemid"]
	item.Sell()
	utils.RespondWithJSON(w, http.StatusOK, "done")
}

func GetPhotos(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	item.ID = mux.Vars(r)["itemid"]
	item.Get()
	utils.RespondWithJSON(w, http.StatusOK, item.Hashtag)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	items := models.GetAllItems()
	utils.RespondWithJSON(w, http.StatusOK, items)
}

func FilterItems(w http.ResponseWriter, r *http.Request) {
	var i models.Item
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &i)
	items := i.Filter()
	utils.RespondWithJSON(w, http.StatusOK, items)
}

func StoreItem(w http.ResponseWriter, r *http.Request) {
	var i models.Item
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &i)
	u1 := uuid.Must(uuid.NewV4())
	i.ID = u1.String()
	i.Store()
	utils.RespondWithJSON(w, http.StatusOK, "Stored")
}
