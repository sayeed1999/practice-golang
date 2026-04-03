package racer_test

import (
	racer "select"
	"testing"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	want := fastURL
	got := racer.Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
