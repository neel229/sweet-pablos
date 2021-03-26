package main

import (
	"fmt"
	"log"

	auth "github.com/neel229/sweet-pablos/internal/auth/api"
	"github.com/neel229/sweet-pablos/util"
)

func main() {
	config, err := util.LoadConfig("./config/auth")
	if err != nil {
		log.Fatalf("cannot read configurations: %v", err)
		return
	}

	server, err := auth.NewServer(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Setting up routes...")
	server.SetRoutes()
	fmt.Println("Starting the server...")
	server.StartServer()
}
