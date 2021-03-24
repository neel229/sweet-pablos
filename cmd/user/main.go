package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	user "github.com/neel229/sweet-pablos/internal/user/api"
	db "github.com/neel229/sweet-pablos/internal/user/db/sqlc"
	"github.com/neel229/sweet-pablos/util"
)

func main() {
	config, err := util.LoadConfig("./config/test")
	if err != nil {
		log.Fatalf("cannot read configurations: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("there was an error creating connection with database: %v", err)
	}

	store := db.NewStore(conn)
	server := user.NewServer(store)
	fmt.Println("Setting up routes...")
	server.SetRoutes()
	fmt.Println("Starting the server...")
	server.StartServer()
}
