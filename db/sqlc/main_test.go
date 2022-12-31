package db

import (
	"database/sql"

	"log"
	"testing"
	_ "github.com/lib/pq"
)

var testQuries *Queries

const (
	dbDriver = "postgres"
	dbconn   = "postgres://postgres:1234@localhost:5432/simplebank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbconn)
	if err != nil {
		log.Println(err.Error())
	}
	testQuries = New(conn)

	m.Run()

}
