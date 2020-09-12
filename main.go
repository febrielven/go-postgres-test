package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/febrielven/go-postgres-test/api/middleware"
)

func main() {
	db := middleware.CreateConnection()

	defer db.Close()

	fmt.Println("Starting server on the port 3999...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
