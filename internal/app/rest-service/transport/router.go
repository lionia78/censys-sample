package transport

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/censys-sample/internal/app/rest-service/controller"
)

// StartServer sets up the router and starts the HTTP server.
func StartServer(controller *controller.KVController, httpAddr string) {
	r := mux.NewRouter()

	// Route mapping using the Controller methods
	r.HandleFunc("/kv/{key}", controller.HandlePut).Methods(http.MethodPost)
	r.HandleFunc("/kv/{key}", controller.HandleGet).Methods(http.MethodGet)
	r.HandleFunc("/kv/{key}", controller.HandleDelete).Methods(http.MethodDelete)

	log.Printf("[StartServer] rest-service listening on %s", httpAddr)
	err := http.ListenAndServe(httpAddr, r)
	log.Fatal(err)
}
