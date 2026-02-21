package main

import (
	"credential-go/database"
	"credential-go/handler"
	"log"
	"net/http"
)

func main() {

	database.Connect()

	http.HandleFunc("/credential", handler.CredentialHandler)
	
	log.Println("Credential service running on :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
