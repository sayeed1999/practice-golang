package concurrency

import "time"

// Here's the setup: a colleague has written a function, CheckWebsites, that checks the status of a list of URLs.

// WebsiteChecker is a function that should take a url and return true/false if valid/invalid
type WebsiteChecker func(url string) bool

func CheckWebsites(ws WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)

	for _, url := range urls {
		// running go test -race . catches race condition on concurrent map writes !!
		// go func() {
		isValid := ws(url)
		result[url] = isValid
		// }()
	}

	time.Sleep(2 * time.Second)

	return result
}
