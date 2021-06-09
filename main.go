package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Phati/demo-load-balancer/db"
	"github.com/Phati/demo-load-balancer/server"
)

func main() {

	db.SetDB()

	router := server.InitRouter()

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
