package concurrency

// Here's the setup: a colleague has written a function, CheckWebsites, that checks the status of a list of URLs.

// WebsiteChecker is a function that should take a url and return true/false if valid/invalid
type WebsiteChecker func(url string) bool

// result is a struct with one string one bool
type result struct {
	url     string
	isValid bool
}

func CheckWebsites(ws WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result)

	for _, url := range urls {
		go func() {
			isValid := ws(url)
			resultChan <- result{url, isValid} // pushing a result instead of concurrent writing
		}()
	}

	for i := 0; i < len(urls); i++ {
		received := <-resultChan
		// fmt.Println("received: ", received)
		results[received.url] = received.isValid // processing the map synchronously with no race condition!
	}

	return results
}
