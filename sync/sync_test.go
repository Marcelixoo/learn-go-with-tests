package sync

import (
	"sync"
	"testing"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
func (c *Counter) Value() int {
	return c.value
}

func TestCounter(t *testing.T) {
	t.Run("it starts at 0", func(t *testing.T) {
		counter := Counter{}

		got := counter.Value()
		want := 0

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounterState(t, &counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		expectedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(expectedCount)

		for i := 0; i < expectedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounterState(t, &counter, expectedCount)
	})
}

func assertCounterState(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf(
			"got %d, want %d",
			got.Value(), want,
		)
	}
}
