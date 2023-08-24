package main

import (
	"fmt"

	concurrency "github.com/Marcelixoo/learn-go-with-tests/concurrency/src"
)

const (
	GOOGLE_DOT_COM_URL  = "https://google.com"
	SPOTIFY_DOT_COM_URL = "https://open.spotify.com/"
	FAKE_WEBSITE_URL    = "waat://furhurterwe.geds"
)

func main() {
	fmt.Println("running concurrency chapter...")

	websites := []string{
		GOOGLE_DOT_COM_URL,
		SPOTIFY_DOT_COM_URL,
		FAKE_WEBSITE_URL,
	}

	wc := func(url string) bool {
		return url != FAKE_WEBSITE_URL
	}

	results := concurrency.CheckWebsites(wc, websites)

	fmt.Printf("Results: %v", results)
}
