package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mgiks/ttyper/internal/server"
)

func main() {
	loadEnvs()

	log.Fatal(http.ListenAndServe(":8000", server.New()))
}

func loadEnvs() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("failed to load environmental variables:", err)
	}
}
