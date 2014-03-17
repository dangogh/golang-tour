package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

func walkImpl(t *tree.Tree, ch chan<- int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch)
	ch <- t.Value
	walkImpl(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan<- int) {
	walkImpl(t, ch)
	close(ch) // walk complete
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
    go Walk(t2, ch2)
	for val1 := range ch1 {
		val2, ok := <-ch2
		if ok && val1 == val2 {
			continue
		}
		return false
	}
    return true
}

func main() {
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("ok 1")
	} else {
		fmt.Println("not ok 1")
	}
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("ok 2")
	} else {
		fmt.Println("not ok 2")
	}

}
