package main

import (
	"log"
	"net/http"

	v0 "github.com/Marcelixoo/learn-go-with-tests/json/v0"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := v0.NewPlayerServer(&InMemoryPlayerStore{})
	log.Fatal(http.ListenAndServe(":50012", server))
}
