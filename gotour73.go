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

type seen struct {
	s map[string]bool
    m sync.Mutex
}

func (s *seen) newToMe(url string) bool {
    defer func() {
        s.s[url] = true
    	s.m.Unlock()
    }()
    s.m.Lock()
    return s.s[url]
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
    if depth <= 0 {
        return
    }

    // make unique
    ch := make(chan string)
    uniq := make(chan string)
    go func() {    
    	var s seen
        for u := range ch {
            if s.newToMe(u) {
            	uniq <- u
            }
        }
    }()
    ch <- url
    
    for url := range uniq {
        go func() {
		    body, urls, err := fetcher.Fetch(url)
		    if err != nil {
        		fmt.Println(err)
		        return
    		}
		    fmt.Printf("found: %s %q\n", url, body)
            for _, u := range urls {
            	ch <- u
            }
        }()
    }

    return
}

func main() {
    Crawl("http://golang.org/", 4, fetcher)
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
