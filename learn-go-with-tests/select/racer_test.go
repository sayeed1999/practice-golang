package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// slowURL := "http://www.facebook.com"
	// fastURL := "http://www.quii.dev"

	// From the book: -
	// httptest.NewServer takes an http.HandlerFunc which we are sending in via an anonymous function.
	// http.HandlerFunc is a type that looks like this: type HandlerFunc func(ResponseWriter, *Request).
	// All it's really saying is it needs a function that takes a ResponseWriter and a Request,
	// which is not too surprising for an HTTP server.
	// There's really no extra magic here, this is also how you would write a real HTTP server in Go.
	// The only difference is we are wrapping it in an httptest.NewServer which makes it easier to use with testing,
	// as it finds an open port to listen on and then you can close it when you're done with your test.

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(3 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL // slowURL
		got, _ := Racer(slowURL, fastURL, 50*time.Millisecond)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 2s", func(t *testing.T) {
		serverA := makeDelayedServer(3 * time.Second)
		serverB := makeDelayedServer(4 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL, 2*time.Second)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}
