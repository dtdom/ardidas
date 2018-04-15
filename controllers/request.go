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

func GetRequest(w http.ResponseWriter, r *http.Request) {
	var request models.Request
	request.ID = mux.Vars(r)["requestid"]
	request.Get()
	utils.RespondWithJSON(w, http.StatusOK, request)
}

func CompleteRequest(w http.ResponseWriter, r *http.Request) {
	var request models.Request
	request.ID = mux.Vars(r)["requestid"]
	request.Complete()
	utils.RespondWithJSON(w, http.StatusOK, "done")
}

func GetRequests(w http.ResponseWriter, r *http.Request) {
	requests := models.GetAllRequests()
	utils.RespondWithJSON(w, http.StatusOK, requests)
}

func StoreRequest(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &req)
	u1 := uuid.Must(uuid.NewV4())
	req.ID = u1.String()
	req.Status = "processing"
	req.Store()
	utils.RespondWithJSON(w, http.StatusOK, "Stored")
}
