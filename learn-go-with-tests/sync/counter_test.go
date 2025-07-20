package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		counter := Counter{}
		wantedCount := 1000

		var wg sync.WaitGroup
		wg.Add(1000)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait() // waits until all 1000 goroutines are released,
		// this helps to not do assertion until all pending operations are finished!

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, counter Counter, want int) {
	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}
