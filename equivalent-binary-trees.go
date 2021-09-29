package main

import "golang.org/x/tour/tree"
import "fmt"



func WalkNode(t *tree.Tree, ch chan int) {
	if(t == nil) {
		return
	}
	WalkNode(t.Left, ch)
	ch <- t.Value
	WalkNode(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkNode(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	
	for v1 := range c1 {
		v2 := <- c2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	k := 1
	
	tree, treeEqual, treeDiff := tree.New(k), tree.New(k), tree.New(2*k)
	
	sameTreesEqual := Same(tree, treeEqual)
	diffTreesEqual := Same(tree, treeDiff)
	fmt.Println(tree, " equals to", treeEqual)
	fmt.Println("Should be true: ", sameTreesEqual)
	fmt.Println(tree, " equals to ", treeDiff)
	fmt.Println("Should be false: ", diffTreesEqual)
}
