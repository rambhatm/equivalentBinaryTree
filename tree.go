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
func getRandomNum(n int) int {
	return (rand.Intn(n) + 1) * 1000
}

//Newtree Returns a random binary tree holding n values k, 2k,...nk
func Newtree(n int, k int) *Tree {
	if n < 1 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	var head = Tree{nil, getRandomNum(k), nil}

	for i := 1; i < n; i++ {
		head.insert(getRandomNum(k))
	}
	return &head
}
