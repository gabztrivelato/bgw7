package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Name struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GreetingsHandler(w http.ResponseWriter, r *http.Request) {
	var name Name
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&name)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	msg := fmt.Sprintf("Hello %s %s", name.FirstName, name.LastName)
	response := map[string]string{"message": msg}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/greetings", GreetingsHandler)
	http.ListenAndServe(":8080", nil)
}
