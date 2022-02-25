package main

import (
	model "backend/models"
	historyService "backend/services/history.service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func CreateOperacionEndpoint(w http.ResponseWriter, req *http.Request) {
	var operacion model.Operacion
	_ = json.NewDecoder(req.Body).Decode(&operacion)
	err := historyService.Create(operacion)
	if err != nil {
		json.NewEncoder(w).Encode("DB ERROR")
	}
	json.NewEncoder(w).Encode(operacion)

}

func GetHistorialEndpoint(w http.ResponseWriter, req *http.Request) {
	historial, err := historyService.Read()

	if err != nil {
		json.NewEncoder(w).Encode("DB ERROR")
	}
	json.NewEncoder(w).Encode(historial)
}

func main() {

	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	ttl := handlers.MaxAge(3600)
	origins := handlers.AllowedOrigins([]string{"*"})
	// endpoints
	router.HandleFunc("/historial", GetHistorialEndpoint).Methods("GET")
	router.HandleFunc("/operacion", CreateOperacionEndpoint).Methods("POST")
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "3000"
	}
	fmt.Println("Backend Started In Port: " + httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, handlers.CORS(headers, credentials, methods, ttl, origins)(router)))
}
