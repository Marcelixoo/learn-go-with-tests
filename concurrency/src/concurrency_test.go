package concurrency_test

import (
	"context"
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

func SendMessageAfter(_ time.Duration, out chan string, msg string, ctx context.Context) {
	// time.Sleep(delayInSeconds * time.Second)
	select {
	case <-ctx.Done():
		return
	default:
		out <- msg
	}
}

type SubRoutine interface {
	Run()
}

func RunMainProcess(subroutines []func(), out chan string, retries int) []string {
	for _, run := range subroutines {
		go run()
	}

	var results []string

	for i := 0; i < retries; i++ {
		select {
		case msg := <-out:
			results = append(results, msg)
		default:
			results = append(results, "neither routine was ready")
		}

		time.Sleep(1 * time.Second)
	}

	return results
}

func RunMainProcessWithoutRetries(subroutines []func(), out chan string, cancel context.CancelFunc) []string {
	for _, run := range subroutines {
		go run()
	}

	defer cancel()

	var results []string
	for msg := range out {
		results = append(results, msg)
	}
	return results
}

func TestSelectStatementForInitialization(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	out := make(chan string)

	t.Run("it waits till some of the channels is ready to communicate", func(t *testing.T) {
		r1 := func() {
			SendMessageAfter(1, out, "received message from subroutine 1", ctx)
		}
		r2 := func() {
			SendMessageAfter(3, out, "received message from subroutine 2", ctx)
		}
		numOfRetries := 5

		got := RunMainProcess([]func(){r1, r2}, out, numOfRetries)

		assertNumberOfMessages(t, got, numOfRetries)

		assertFirstMessage(t, got, "neither routine was ready")
		assertLastMessage(t, got, "received message from subroutine 2")
	})

	t.Run("it drains the channel til the end", func(t *testing.T) {
		after1Sec := func() {
			SendMessageAfter(1, out, "received message from subroutine 1", ctx)
		}

		got := RunMainProcessWithoutRetries([]func(){
			after1Sec,
			after1Sec,
			after1Sec,
		}, out, cancel)

		defer cancel()

		assertNumberOfMessages(t, got, 2)
	})
}

func assertNumberOfMessages(t *testing.T, got []string, want int) {
	t.Helper()
	numOfMessages := len(got)
	if numOfMessages != want {
		t.Errorf("expected %d messages got %d", want, numOfMessages)
	}
}

func assertFirstMessage(t *testing.T, got []string, match string) {
	t.Helper()

	position := 0
	first := got[position]

	if first != match {
		t.Errorf("unexpected results; want %q got %q at position %d", match, first, position)
	}
}

func assertLastMessage(t *testing.T, got []string, match string) {
	t.Helper()

	position := len(got) - 1
	last := got[position]

	if last != match {
		t.Errorf("unexpected results; want %q got %q at position %d", match, last, position)
	}
}
