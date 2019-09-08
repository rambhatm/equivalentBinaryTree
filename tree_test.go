package main

import (
	"fmt"
	"testing"
)

func TestTreeAdd(t *testing.T) {
	aRandomTree := Newtree(100, 5)

	if aRandomTree == nil {
		t.Errorf("Unable to create a new random tree")
	}
	fmt.Println(aRandomTree)
}

func TestWalkTree(t *testing.T) {
	t100 := Newtree(100, 10)
	var i int

	ch := make(chan int)
	go Walk(t100, ch)
	for i = range ch {
		fmt.Println(i)
	}
	if i != 100 {
		t.Errorf("Walk not returning all values %d/100", i)
	}
}
