package main

import (
	"log"
	"net/http"

	"github.com/mgiks/typo-typer/server"
)

func main() {
	http.HandleFunc("GET /texts", server.GETTextHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
