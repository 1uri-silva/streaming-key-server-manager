package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	dbEnvUser := os.Getenv("POSTGRES_USER")
	dbEnvPass := os.Getenv("POSTGRES_PASS")
	dbEnvPort := os.Getenv("POSTGRES_PORT")
	dbEnvHost := os.Getenv("POSTGRES_HOST")
	dbEnvDb := os.Getenv("POSTGRES_DB")

	port, err := strconv.Atoi(dbEnvPort)

	if err != nil {
		log.Fatalf("Error convert port string in port number")
	}

	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbEnvHost, port, dbEnvUser, dbEnvPass, dbEnvDb)

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}
