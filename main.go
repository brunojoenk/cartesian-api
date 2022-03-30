package main

import (
	"log"
	"os"

	"github.com/brunojoenk/cartesian-api/internal/servers"
)

func main() {
	server := servers.NewServer()

	os.Setenv("PATH_DATA_POINTS", "data/points.json")

	log.Printf("Server started at address: %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
