package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver
)

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
