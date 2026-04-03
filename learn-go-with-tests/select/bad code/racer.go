package racer

import (
	"net/http"
	"time"
)

// You have been asked to make a function called WebsiteRacer which takes two URLs
// and "races" them by hitting them with an HTTP GET and returning the URL which returned first.
// If none of them return within 10 seconds then it should return an error.

func Racer(url1, url2 string) (winner string) {

	timeA := measureResponseTime(url1)
	timeB := measureResponseTime(url2)

	if timeA < timeB {
		return url1
	}

	return url2
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
