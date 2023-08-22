package greeting_test

import (
	"testing"

	"github.com/Marcelixoo/learn-go-with-tests/meta/greeting"
)

func TestHello(t *testing.T) {
	got := greeting.Hello("Chris", "es")
	want := "Hola, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
