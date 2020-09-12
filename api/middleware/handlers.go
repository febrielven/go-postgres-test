package middleware

import (
	"database/sql" // package db sql
	"fmt"
	"log"
	"os" // used to read the environment variable

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
)

// Response format
type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// CreateConnection with postgress
func CreateConnection() *sql.DB {
	//load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	//Cek Connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connection")

	//Return the connection
	return db
}


