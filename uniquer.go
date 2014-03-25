package main

import "fmt"

type empty struct{}

type uniquer struct {
	in chan string
	out chan bool
	m map[string]empty
}

func (u *uniquer) isUniq(s string) {
	in <- s
	return <-out
}

func (u *uniquer) init() {
	go func() {
		for s := range u.in {
			seenit := false
			if _, seenit = u[s]; !seenit {
				// not found -- add it
				u[s] = empty{}
			}
			u.out <- !seenit
		}
	}()
}


func main() {
	var u uniquer{make(chan string), make(chan empty), {}}
	u.init()
	urls := []string{"a", "b", "c", "b", "a", "d"}
	for _, s := range urls {
		fmt.Println(s, " unique? ", u.isUniq(s))
	}
}
