package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func crawl(crawlCnOutput chan CrawlRequest, crawlCnInput chan CrawlRequest, wg *sync.WaitGroup, fetcher Fetcher) {
	// fmt.Println("Starting goroutine crawl")
	// for r := range crawlCn {
	visitedUrls := make(map[string]string)
	for {
		r, ok := <-crawlCnOutput
		if !ok {
			return
		}

		if _, ok := visitedUrls[r.url]; ok {
			wg.Done()
			continue
		} else {
			visitedUrls[r.url] = "hi"
		}

		// fmt.Printf("received crawl task %s (depth:%d)\n", r.url, r.depth)

		if r.depth > 0 {
			go func() {
				body, urls, err := fetcher.Fetch(r.url)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("found: %s %q\n", r.url, body)
					for _, u := range urls {
						wg.Add(1)
						crawlCnInput <- CrawlRequest{u, r.depth - 1}
					}
				}
				wg.Done()
			}()
		}
	}
}

type CrawlRequest struct {
	url   string
	depth int
}

func deduplicateMessages(inputCn chan CrawlRequest) (outputCn chan CrawlRequest) {
	// fmt.Println("Starting goroutine deduplicateMessage")
	returnCn := make(chan CrawlRequest)

	go func() {
		for input := range inputCn {
			// fmt.Printf("received request %s (depth:%d)\n", input.url, input.depth)
			returnCn <- input
			// fmt.Printf("sent request %s (depth:%d)\n", input.url, input.depth)
		}
		close(returnCn)
	}()

	return returnCn
}

func WebCrawlerMain() {

	crawlCnInput := make(chan CrawlRequest)
	crawlCnOutput := deduplicateMessages(crawlCnInput)
	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		crawlCnInput <- CrawlRequest{"https://golang.org/", 4}

		wg.Wait()
		close(crawlCnInput)
	}()

	crawl(crawlCnOutput, crawlCnInput, &wg, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
