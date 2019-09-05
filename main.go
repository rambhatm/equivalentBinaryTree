package main

type Tree struct {
	left  *Tree
	value int
	right *Tree
}

//Inserts integer into sorted position in the tree
func (t *Tree) insert(x int) {
	if x < t.value {
		t.left.insert(x)
	} else {
		t.right.insert(x)
	}

}

func main() {
}
