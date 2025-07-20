package racer

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string, timeout time.Duration) (winner string, err error) {

	// aDuration := measureResponseTime(a)
	// bDuration := measureResponseTime(b)

	// if aDuration > bDuration {
	// 	return b
	// }

	// return a

	select { // select allows waiting on multiple channels simultaneously
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): // works like default clause with a timeout for select
		return "", errors.New(fmt.Sprintf("operation timeout: exceeded %v", timeout))
	}
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	duration := time.Since(start)
// 	return duration
// }

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
