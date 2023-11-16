package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver  = "postgres"
	dbSourced = "postgresql://devmarcos:karybe021313@go-finance.cayl5ldqzc4i.us-east-1.rds.amazonaws.com:5432/postgres?sslmode=require"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSourced)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
