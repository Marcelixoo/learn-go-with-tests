package concurrency

type WebsiteChecker func(string) bool

type Result struct {
	URL   string
	Valid bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	channel := make(chan Result)

	for _, url := range urls {
		/*
			Anonymous functions keep lexical scope–all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.
		*/
		go func(url string) {
			// send result to channel
			channel <- Result{url, wc(url)}
		}(url)
	}

	for range urls {
		// receive result from channel
		r := <-channel
		results[r.URL] = r.Valid
	}

	return results
}
