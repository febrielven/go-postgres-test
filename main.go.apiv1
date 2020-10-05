package main

import (
	"fmt"
	"go-postgres-test/api/routes"
	"log"
	"net/http"
)

func main() {
	// db := middleware.CreateConnection()
	// defer db.Close()
	r := routes.Router()

	fmt.Println("Starting server on the port 8082...")
	log.Fatal(http.ListenAndServe(":8082", r))
}
