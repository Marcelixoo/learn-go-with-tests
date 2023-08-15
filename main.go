package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running chapter \"JSON, routing and embedding\"")

	server := &PlayerServer{NewInMemoryPlayerStore()}

	err := http.ListenAndServe(":5000", server)
	log.Fatal(err)
}
