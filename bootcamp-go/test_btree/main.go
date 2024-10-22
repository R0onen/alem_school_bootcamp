package main

import (
	"bootcamp/btree"
	"fmt"
)

func main() {
	tree := btree.NewBTree()
	tree.ReplaceOrInsert(50)
	tree.ReplaceOrInsert(30)
	tree.ReplaceOrInsert(70)

	printTree(tree.Root) // 30 50 70
	/* Btree Visualization:
	      50
	    /    \
	   30    70
	*/

	tree.Clear()
	printTree(tree.Root) // Outputs nothing as the tree is now empty
}

func printTree(node *btree.BTreeNode) {
	if node == nil {
		return
	}
	printTree(node.Left)
	fmt.Println(node.Value)
	printTree(node.Right)
}
