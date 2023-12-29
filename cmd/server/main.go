package main

import (
	"log"

	"github.com/daniel-pk-harrison/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
