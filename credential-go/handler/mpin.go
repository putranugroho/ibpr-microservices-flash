package handler

import (
	"credential-go/database"
	"credential-go/service"
	"encoding/json"
	"net/http"
)

func CredentialHandler(w http.ResponseWriter, r *http.Request) {

	var req service.CredentialRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    "001",
			"status":  "error",
			"message": "Invalid request body",
			"data":    nil,
		})
		return
	}

	result, err := service.CheckCredential(database.DB, req)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    "099",
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	json.NewEncoder(w).Encode(result)
}
