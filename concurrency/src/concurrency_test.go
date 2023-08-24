package concurrency_test

import (
	"reflect"
	"testing"
	"time"

	concurrency "github.com/Marcelixoo/learn-go-with-tests/concurrency/src"
)

const (
	GOOGLE_DOT_COM_URL  = "https://google.com"
	SPOTIFY_DOT_COM_URL = "https://open.spotify.com/"
	FAKE_WEBSITE_URL    = "waat://furhurterwe.geds"
)

func mockWebsiteChecker(url string) bool {
	return url != FAKE_WEBSITE_URL
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		GOOGLE_DOT_COM_URL,
		SPOTIFY_DOT_COM_URL,
		FAKE_WEBSITE_URL,
	}

	want := map[string]bool{
		GOOGLE_DOT_COM_URL:  true,
		SPOTIFY_DOT_COM_URL: true,
		FAKE_WEBSITE_URL:    false,
	}

	got := concurrency.CheckWebsites(mockWebsiteChecker, websites)

	time.Sleep(2 * time.Second)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		concurrency.CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
