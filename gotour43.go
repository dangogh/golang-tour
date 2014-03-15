package main

import (
    "strings"
    "code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
    counts := make(map[string]int)
    for _, w := range(strings.Fields(s)) {
    	counts[w]++
    }
    return counts
}

func main() {
    wc.Test(WordCount)
}
