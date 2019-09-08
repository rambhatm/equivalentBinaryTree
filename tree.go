package main

import (
	"math/rand"
	"time"
)

type Tree struct {
	left  *Tree
	value int
	right *Tree
}

//Inserts integer into sorted position in the tree
func (t *Tree) insert(x int) {

	if x < t.value {
		if t.left == nil {
			t.left = &Tree{value: x}
			return
		}
		t.left.insert(x)
	} else {
		if t.right == nil {
			t.right = &Tree{value: x}
			return
		}
		t.right.insert(x)
	}

}

//Returns a string representation of the tree
func (t *Tree) String() string {
	return "abcd"
}

//generates random num (1-n)K
func generateValue(n int) int {
	return (rand.Intn(n) + 1) * 1000
}

//Newtree Returns a random binary tree holding n values k, 2k,...nk
func Newtree(n int, k int) *Tree {
	if n < 1 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	var head = Tree{value: generateValue(k)}

	for i := 1; i < n; i++ {
		head.insert(generateValue(k))
	}
	return &head
}
func walkTree(t *Tree, ch chan int) {
	if t.left != nil {
		walkTree(t.left, ch)
	}
	ch <- t.value
	if t.right != nil {
		walkTree(t.right, ch)
	}

}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	walkTree(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for val := range ch1 {
		if val != <-ch2 {
			return false
		}
	}
	return true
}
