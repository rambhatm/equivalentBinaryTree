package main

import (
	"fmt"
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
	if t == nil {
		head := new(Tree)
		head.value = x
		fmt.Printf("\n New tree head %p", head)
		t = head
	}
	/*
		if x < t.value {
			t.left.insert(x)
		} else {
			t.right.insert(x)
		}

		leaf := new(Tree)
		leaf.value = x
	*/
}

func getRandomNum(n int) int {
	return (rand.Intn(n) + 1) * 1000
}

//Returns a random binary tree holding values k, 2k,...nk
func Newtree(n int, k int) *Tree {
	if n < 1 {
		return nil
	}
	var head = Tree{nil, getRandomNum(k), nil}
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < n; i++ {
		head.insert(getRandomNum(k))
	}
	fmt.Printf("\n New tree head2 %p", &head)
	return &head
}
