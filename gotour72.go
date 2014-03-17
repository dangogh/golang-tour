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
	return false
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for ii := range ch {
		fmt.Println(ii)
	}
}
