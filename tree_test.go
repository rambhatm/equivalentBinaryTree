package eqbinarytree

import (
	"fmt"
	"testing"
)

func TestTreeAdd(t *testing.T) {
	aRandomTree := Newtree(10, 5)

	if aRandomTree == nil {
		t.Errorf("Unable to create a new random tree")
	}
	fmt.Println(aRandomTree)
}

func TestSame(t *testing.T) {
	t100 := Newtree(100, 10)
	if Same(t100, t100) != true {
		t.Errorf("Test Same not working")
	}
}
