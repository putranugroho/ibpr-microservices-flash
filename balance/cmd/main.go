package main

import (
	"balance/database"
	"balance/handler"
	"log"
	"net/http"
)

func main() {

	database.Connect()

	http.HandleFunc("/saldo", handler.BalanceHandler)
	
	log.Println("Balance service running on :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
