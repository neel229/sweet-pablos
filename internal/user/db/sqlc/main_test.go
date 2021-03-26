package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/neel229/sweet-pablos/util"
)

// testQueries will be used for testing
// sql queries
var testQueries *Queries
var testDB *sql.DB

// Entry test where we setup our DB connection
func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../../../../config/user")
	if err != nil {
		log.Fatalf("error opening the config file: %v", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
