package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rafaelfcads/file-api/route"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}
	router := route.NewRouter()

	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
