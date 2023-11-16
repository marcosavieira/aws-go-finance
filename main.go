package main

import (
	"database/sql"
	"log"
	"os"

	"https://github.com/marcosavieira/aws-go-finance/api"
	db "https://github.com/marcosavieira/aws-go-finance/db/sqlc"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading godotenv: ", err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbSourced := os.Getenv("DB_SOURCED")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	conn, err := sql.Open(dbDriver, dbSourced)
	if err != nil {
		log.Fatal("cannot connect to database", err)
		return
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start api", err)
		return
	}
}
