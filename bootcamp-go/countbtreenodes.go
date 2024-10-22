package bootcamp

import (
	"bootcamp/btree"
)

func CountBtreeNodes(b *btree.BTree) int {
	return countBtreeNodesRecursive(b.Root)
}

func countBtreeNodesRecursive(node *btree.BTreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + countBtreeNodesRecursive(node.Left) + countBtreeNodesRecursive(node.Right)
}
