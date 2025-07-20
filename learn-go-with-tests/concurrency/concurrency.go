package concurrency

// This is a func that takes a URL as string and a bool as alive or not
type WebsiteChecker func(string) bool

// Using Dependency Injection (DI) has allowed us to test this func without making real http calls!
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		ok := wc(url)
		results[url] = ok
	}

	return results
}
