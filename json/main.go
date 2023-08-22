package main

import (
	"log"
	"net/http"

	v0 "github.com/Marcelixoo/learn-go-with-tests/json/v0"
)

func main() {
	store := v0.NewInMemoryPlayerStore()
	server := v0.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":50012", server))
}
