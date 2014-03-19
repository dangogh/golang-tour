package main

import (
    "fmt"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

type empty struct {}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan<- string, done chan<- empty) {
    defer func() {
        done <- empty{}
    }()
    if depth <= 0 {
        return
    }

    body, urls, err := fetcher.Fetch(url)
    if err != nil {
   		fmt.Println(err)
        return
	}
    ch <- url
    fmt.Printf("found: %s %q\n", url, body)

    waitforN := make(chan empty)
    for _, u := range urls {
      go Crawl(u, depth-1, fetcher, ch, waitforN)
    }
    // wait for all to finish before declaring done
    for _, _ = range urls {
    	<- waitforN
    }
    done <- empty{}
    return
}

func main() {
    // make unique
    ch := make(chan string)
    done := make(chan empty)
    Crawl("http://golang.org/", 4, fetcher, ch, done)
    
    uniq := make(chan string)
    var seen map[string]bool
    go func() {
        for url := range ch {
            if _, ok := seen[url]; !ok {
                continue
            }
            seen[url] = true
            uniq <- url
        }
	    // wait for top one to finish
      	<-done
        close(uniq)
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
