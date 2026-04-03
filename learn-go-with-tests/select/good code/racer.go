package racer

import (
	"fmt"
	"net/http"
	"time"
)

// You have been asked to make a function called WebsiteRacer which takes two URLs
// and "races" them by hitting them with an HTTP GET and returning the URL which returned first.
// If none of them return within 10 seconds then it should return an error.

// This time we don't wait for both GET to complete, rather choose winner who reaches first without calculating time!

func Racer(url1, url2 string) (winner string, err error) {

	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(9 * time.Second):
		return "", fmt.Errorf("operation timed out: 10s")
	}
}

// ping creates a channel of unit value and returns it
func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch) // signals when GET request is finished!
	}()

	return ch
}
