package concurrency

import "time"

// WebsiteChecker checks a url, returning a bool.
type WebsiteChecker func(string) bool

// CheckWebsites throws a race condition
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)
	return results
}
