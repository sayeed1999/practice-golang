package racer_test

import (
	"net/http"
	"net/http/httptest"
	racer "select"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speed of servers, returns the fastest one", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(1 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, _ := racer.Racer(slowServer.URL, fastServer.URL)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("returns an error if the server doesn't respond within 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Second)
		fastServer := makeDelayedServer(12 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		// we expect an error since both server should take > 10s
		_, err := racer.Racer(fastServer.URL, slowServer.URL)

		if err == nil {
			t.Error("expected error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
