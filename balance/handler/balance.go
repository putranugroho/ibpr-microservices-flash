package handler

import (
	"balance/database"
	"balance/service"
	"encoding/json"
	"net/http"
)

func BalanceHandler(w http.ResponseWriter, r *http.Request) {

	var req service.BalanceRequest
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

	result, err := service.CheckBalance(database.DB, req)
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
