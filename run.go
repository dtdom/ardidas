package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dtdom/ardidas/controllers"
	"github.com/rs/cors"

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
	r.Router.HandleFunc("/request/{requestid}/done", controllers.CompleteRequest).Methods("GET")
	r.Router.HandleFunc("/request/{requestid}", controllers.GetRequest).Methods("GET")
	r.Router.HandleFunc("/request", controllers.GetRequests).Methods("GET")
	r.Router.HandleFunc("/request", controllers.StoreRequest).Methods("POST")
}

func (r *Runner) Run(addr string) {
	fmt.Println("listening on port ", addr)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r.Router)
	log.Fatal(http.ListenAndServe(addr, handler))
}
