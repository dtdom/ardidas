package main

import (
	"ardidas/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Runner struct {
	Router *mux.Router
}

func (r *Runner) Initialize() {
	r.Router = mux.NewRouter()
	r.initializeRoutes()
}

func (r *Runner) initializeRoutes() {
	r.Router.HandleFunc("/item/{itemid}", controllers.GetItem).Methods("GET")
	r.Router.HandleFunc("/item/{itemid}/sell", controllers.SellItem).Methods("GET")
	r.Router.HandleFunc("/item/filter", controllers.FilterItems).Methods("POST")
	r.Router.HandleFunc("/item", controllers.StoreItem).Methods("POST")
	r.Router.HandleFunc("/item", controllers.GetItems).Methods("GET")
	r.Router.HandleFunc("/item/{itemid}/photos", controllers.GetPhotos).Methods("GET")
	r.Router.HandleFunc("/request/{requestid}", controllers.GetRequest).Methods("GET")
	r.Router.HandleFunc("/request", controllers.GetRequests).Methods("GET")
	r.Router.HandleFunc("/request", controllers.StoreRequest).Methods("POST")
	r.Router.HandleFunc("/request/{requestid}", controllers.CompleteRequest).Methods("PUT")

}

func (r *Runner) Run(addr string) {
	fmt.Println("listening on port ", addr)
	log.Fatal(http.ListenAndServe(addr, r.Router))

}
