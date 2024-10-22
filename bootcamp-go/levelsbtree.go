package bootcamp

import (
	"bootcamp/btree"
)

func LevelsBtree(b *btree.BTree) int {
	return levelsBtreeRecursive(b.Root)
}

func levelsBtreeRecursive(node *btree.BTreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if node == nil {
		return 0
	}
	return 1 + max(levelsBtreeRecursive(node.Left), levelsBtreeRecursive(node.Right))
}
