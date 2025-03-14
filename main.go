package main

import (
	"encoding/json"
	"go-hexagonal-api/core"
	"go-hexagonal-api/infrastructure"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Repo          core.DataRepository
	MessageSender core.MessageSender
}

func NewAPI(repo core.DataRepository, sender core.MessageSender) *API {
	return &API{
		Repo:          repo,
		MessageSender: sender,
	}
}

func (api *API) HandlePostData(w http.ResponseWriter, r *http.Request) {
	var data core.Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Guardar en la base de datos
	if err := api.Repo.SaveData(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar mensaje a MQTT
	if err := api.MessageSender.SendMessage(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func main() {
	repo, err := infrastructure.NewMySQLRepository("root:root@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	sender, err := infrastructure.NewMQTTSender("tcp://184.72.149.16:1883", "clientID")
	if err != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", err)
	}

	api := NewAPI(repo, sender)

	r := mux.NewRouter()
	r.HandleFunc("/data", api.HandlePostData).Methods("POST")

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
