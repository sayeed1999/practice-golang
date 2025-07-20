package concurrency

// This is a func that takes a URL as string and a bool as alive or not
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// Note 1: -
// Using Dependency Injection (DI) has allowed us to test this func without making real http calls!

// Note 2: -
// Using Channels & go routine in Go, we parallely run as many website calls using different processes.

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() { // this pulls up a separate go-routine
			ok := wc(url)
			resultChannel <- result{url, ok}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
