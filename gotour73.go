package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type empty struct{}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan<- string, uniq <-chan string, done chan<- empty) {
	defer func() {
		fmt.Println("Level ", depth, " done")
		done <- empty{}
	}()

	if depth <= 0 {
		return
	}

	fmt.Println("Starting Level ", depth)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	ch <- url

	waitforN := make(chan empty)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, ch, waitforN)
	}
	// wait for all to finish before declaring done
	fmt.Println("waiting for ", len(urls), " to finish at level ", depth)
	for _, _ = range urls {
		<-waitforN
	}
	fmt.Println("done waiting for ", len(urls), " to finish at level ", depth)
	return
}

func main() {
	// make unique
	ch := make(chan string)
	uniq := make(chan string)
	done := make(chan empty)
	go Crawl("http://golang.org/", 4, fetcher, ch, uniq, done)

	var seen map[string]bool
	go func() {
		for {
			select {
			case 
			if _, ok := seen[url]; !ok {
				continue
			}
			seen[url] = true
			uniq <- url
		}
		// wait for top one to finish
	}()

	fmt.Println("waiting for done")
	<-done

	for url := range uniq {
		fmt.Println("Got ", url)
	}
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
