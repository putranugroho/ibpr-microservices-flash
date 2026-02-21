package main

import (
	"mpin-go/database"
	"mpin-go/handler"
	"log"
	"net/http"
)

func main() {

	database.Connect()

	http.HandleFunc("/mpin", handler.MpinHandler)
	
	log.Println("Mpin service running on :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
