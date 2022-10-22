package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type tmpResponse struct {
	val string
	num int
}

func GetUserBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := tmpResponse{"test", 10}

	jsonResponse, jsonError := json.Marshal(resp)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	w.Write(jsonResponse)
}