package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/Marcelixoo/learn-go-with-tests/json/v1"
)

func main() {
	fmt.Println("Running chapter \"JSON, routing and embedding\"")

	server := v1.NewServer(v1.NewInMemoryPlayerStore())

	err := http.ListenAndServe(":50012", server)
	log.Fatal(err)
}
