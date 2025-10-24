package controller

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/censys-sample/internal/app/rest-service/model"
)

// KVController holds a reference to the KVModel interface.
type KVController struct {
	KVService model.KVModel
}

type putBody struct {
	Value string `json:"value"`
}

// HandlePut handles POST /kv/{key}
func (c *KVController) HandlePut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad request: could not read body", http.StatusBadRequest)
		return
	}
	var body putBody
	if err = json.Unmarshal(b, &body); err != nil {
		http.Error(w, "bad json format", http.StatusBadRequest)
		return
	}

	// Delegate to the Model
	if err = c.KVService.Put(context.Background(), key, body.Value); err != nil {
		log.Printf("[controller.HandlePut] kv-service failed to put, error='%v'", err)
		http.Error(w, "kv backend error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// HandleGet handles GET /kv/{key}
func (c *KVController) HandleGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	// Delegate to the Model
	value, found, err := c.KVService.Get(context.Background(), key)
	if err != nil {
		log.Printf("[controller.HandleGet] kv-service failed to get, error='%v'", err)
		http.Error(w, "kv backend error", http.StatusInternalServerError)
		return
	}
	if !found {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	// Respond with JSON (View Equivalent)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}

// HandleDelete handles DELETE /kv/{key}
func (c *KVController) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	// Delegate to the Model
	found, err := c.KVService.Delete(context.Background(), key)
	if err != nil {
		log.Printf("[controller.HandleDelete] kv-service failed to delete, error='%v'", err)
		http.Error(w, "kv backend error", http.StatusInternalServerError)
		return
	}

	if !found {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}
