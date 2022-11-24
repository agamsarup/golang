package main

import (
	"fmt"

	"github.com/agamsarup/golang/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Value != 0 {
		ch <- t.Value
	}

	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	i := 0
	for i < 10 {
		i++
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	i := 0
	for i < 10 {
		i++
		fmt.Printf("%d ", <-ch)
	}

	fmt.Print(Same(tree.New(1), tree.New(1)))
}
