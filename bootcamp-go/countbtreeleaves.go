package bootcamp

import (
	"bootcamp/btree"
)

func CountBtreeLeaves(b *btree.BTree) int {
	return countBtreeLeavesRecursive(b.Root)
}

func countBtreeLeavesRecursive(node *btree.BTreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == node.Right {
		return 1
	}
	return countBtreeLeavesRecursive(node.Left) + countBtreeLeavesRecursive(node.Right)
}
