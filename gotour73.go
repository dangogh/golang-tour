package main

import (
	"fmt"
)

type Uniquer interface {
	IsUniq(s string) bool
}

type uniquer struct {
	in chan string
	out chan empty
	m map[string]empty
}

type Fetcher interface {
	Uniquer
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type empty struct{}

func (u *Uniquer) IsUniq(s string) bool {
	u.in <- s
	return <-u.out
}

func (u *Uniquer) makeUniquer() func(s string) {
	in := make(chan string)
	out := make(chan empty)
	seen := make(map[string]empty)
	
	go func() {
		for s := range ch {
			seenit := false
			if _, seenit = seen[s]; !seenit {
				// not found -- add it
				seen[s] = empty{}
			}
			uniq <- !seenit
		}
	}()
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
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

	return
}

func main() {

	fetcher.makeUniquer()
	go Crawl("http://golang.org/", 4, fetcher)

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

	for url := range uniq {
		fmt.Println("Got ", url)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
	u uniquer
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
