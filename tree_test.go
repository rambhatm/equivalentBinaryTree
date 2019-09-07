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
