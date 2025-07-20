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

	slowServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

	fastServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
