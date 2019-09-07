package main

import (
	"testing"
)

func TestTreeAdd(t *testing.T) {
	aRandomTree := Newtree(1, 1)

	if aRandomTree == nil {
		t.Errorf("Unable to create a new random tree")
	}
}
